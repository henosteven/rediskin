package main

import (
    "net/http"
    "fmt"
    "reflect"
    "context"
    "os"
    "os/exec"
    "os/signal"
    "syscall"
    "log"
    "strconv"
    "sync"
)

func index(w http.ResponseWriter, r *http.Request) {    
    pid := int64(syscall.Getpid())
    fmt.Fprintf(w, "hello server~ webotherway" + strconv.FormatInt(pid, 10))
    fmt.Println(r.Context())
}

func handleSigHook(sig os.Signal) {
    log.Println("ready to start new web server")
    server.Close()
    cmd := exec.Command("go", "run", "webotherway.go")
    err := cmd.Start()
    if err != nil {
        log.Fatal("fatal to run go run webotherway.go")
    } else {
        log.Fatal("good~ tell the parent to exit")
        wg.Done()
    }
}

var (
    server http.Server
    wg sync.WaitGroup
)

func main() {
    
    wg.Add(1)
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
    go func() {
        sig := <-ch
        handleSigHook(sig)
    }()

    serveMux := http.NewServeMux()
    serveMux.HandleFunc("/", index)
    fmt.Println(reflect.TypeOf(serveMux)) //*http.ServeMux
    
    newServeMux := middleWare(serveMux)
    server = http.Server{Addr: ":8080", Handler:newServeMux}
    go server.ListenAndServe()

    wg.Wait()
    
}

func middleWare(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       pctx := r.Context()
       //cctx := context.WithValue(pctx, []int{1,2}, "henosteven") //key is not comparable panic
       //没有任何意义，就是测试而已,可以用来做个统一的参数处理，例如放入一个traceid
       cctx := context.WithValue(pctx, [2]int{1,2}, "henosteven") //key is not comparable panic
       next.ServeHTTP(w, r.WithContext(cctx)) 
    })
}
