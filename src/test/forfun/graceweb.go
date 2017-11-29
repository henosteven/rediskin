package main

import (
    "log"
    "net/http"
    "net"
    "os"
    "os/signal"
    "syscall"
    "fmt"
    "runtime"
    "sync"
    "os/exec"
    "flag"
    "strconv"
)

var (
    httpWg sync.WaitGroup
    netListener *gracefulListener
)

type gracefulListener struct {
    net.Listener
    stop    chan error
    stopped bool
}

func (gl *gracefulListener) Accept() (c net.Conn, err error) {
    c, err = gl.Listener.Accept()
    if err != nil {
        return
    }

    c = gracefulConn{Conn: c}

    httpWg.Add(1)
    return
}

func newGracefulListener(l net.Listener) (gl *gracefulListener) {
    gl = &gracefulListener{Listener: l, stop: make(chan error)}
    go func() {
        _ = <-gl.stop
        gl.stopped = true
        gl.stop <- gl.Listener.Close()
    }()
    return
}

func (gl *gracefulListener) Close() error {
    if gl.stopped {
        return syscall.EINVAL
    }
    gl.stop <- nil
    return <-gl.stop
}

func (gl *gracefulListener) File() *os.File {
    tl := gl.Listener.(*net.TCPListener)
    fl, _ := tl.File()
    return fl
}

type gracefulConn struct {
    net.Conn
}

func (w gracefulConn) Close() error {
    httpWg.Done()
    return w.Conn.Close()
}

func main() {
    log.Print(syscall.Getpid())
    runtime.GOMAXPROCS(runtime.NumCPU())

    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGTERM)
    signal.Notify(ch, syscall.SIGINT)
    go func() {
        <-ch
        file := netListener.File() // this returns a Dup()
        args := []string{
            "-graceful"}

        cmd := exec.Command("./graceweb", args...)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.ExtraFiles = []*os.File{file}

        err := cmd.Start()
        if err != nil {
            log.Fatalf("gracefulRestart: Failed to launch, error: %v", err)
        }
    } ()
    
    servermux := http.NewServeMux()
    servermux.HandleFunc("/", index)
  
    server := &http.Server{Addr: "0.0.0.0:8080", Handler:servermux}

    var gracefulChild bool
    var l net.Listener
    var err error

    flag.BoolVar(&gracefulChild, "graceful", false, "listen on fd open 3 (internal use only)")
    flag.Parse()

    if gracefulChild {
        log.Print("main: Listening to existing file descriptor 3.")
        f := os.NewFile(3, "")
        l, err = net.FileListener(f)
    } else {
        log.Print("main: Listening on a new file descriptor.")
        l, err = net.Listen("tcp", server.Addr)
    } 

    if err != nil {
        log.Panic("failed to listen on:", server.Addr)
    }

    if gracefulChild {
        parent := syscall.Getppid()
        log.Printf("main: Killing parent pid: %v", parent)
        syscall.Kill(parent, syscall.SIGTERM)
    }    

    netListener = newGracefulListener(l)
    server.Serve(netListener)
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello~graceful~server" + strconv.FormatInt(int64(syscall.Getpid()), 10))
}
