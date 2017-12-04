package main

import (
    "fmt"
    "net"
    "os"
    "reflect"
)

type Mlistener struct {
    net.Listener
}

func newMlistener(ln net.Listener) *Mlistener {
    ml := &Mlistener{ln}
    return ml
}

func (ml *Mlistener) Accept() (conn net.Conn, err error) {
    fmt.Println("we can do something before real accept")
    conn, err = ml.Listener.Accept()
    return 
}

func (ml *Mlistener) Close() error {
    fmt.Println("yes - we receive exit signal")
    err := ml.Listener.Close()
    return err
}

func main() {
    ln , err := net.Listen("tcp", ":8090")
    fmt.Println(reflect.TypeOf(ln)) //*net.TCPListener
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    ml := newMlistener(ln)

    defer func() {
        ml.Close()
    }()

    conn, _ := ml.Accept()
    fmt.Println(reflect.TypeOf(conn))
}
