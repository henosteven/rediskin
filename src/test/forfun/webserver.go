/**
 * 如何优雅关闭http server
 * @todo
 */
package main

import (
    "fmt"
    "net/http"
    "context"
    "sync"
    "time"
    "os"
    "os/signal"
    "syscall"
)

var (
    wg sync.WaitGroup
)

func main() {
    
    ctx, cancel := context.WithCancel(context.Background())

    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGTERM)
    signal.Notify(ch, os.Interrupt)
    go func() {
        <-ch
        fmt.Println("bye~bye")
        cancel()
        time.Sleep(time.Second * 3)
        os.Exit(1)
    } ()

    wg.Add(1)
    go startServer(ctx)
    wg.Wait()
}

func startServer(ctx context.Context) {
    defer wg.Done()
    fmt.Println("ready to start  webserver")
    http.HandleFunc("/", Index)
    err := http.ListenAndServe(":6366", nil)
    if err != nil {
        fmt.Println("failed to start web server:", err)
    }

    select {
        case <-ctx.Done() :
            fmt.Println("should be exit!")
    }
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello server~")
}
