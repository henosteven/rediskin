package main

import (
    "fmt"
    "context"
    "sync"
    "time"
)

var wg sync.WaitGroup

func main() {
    ctx := context.Background()
    cctx := context.WithValue(ctx, "key-user", "value-heno")
    wg.Add(1)
    go child1(cctx)
    wg.Wait()
}

func child1(ctx context.Context) {
    defer wg.Done() 
    for {
        time.Sleep(time.Second * 1)
        fmt.Println(ctx)
    }
}
