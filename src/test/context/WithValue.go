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
    //func WithValue(parent Context, key, val interface{}) Context
    //注意这里的interface{}, key 和 val 都可以是任意类型的值
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
