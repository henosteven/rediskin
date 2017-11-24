package main

import (
    "context"
    "log"
    "os"
    "time"
    "fmt"
)

var logg *log.Logger

func someHandler() {
    pctx := context.Background()
    ctx, cancel := context.WithCancel(pctx)
    go doStuff(ctx)
    go doRequest(ctx)
    time.Sleep(10 * time.Second)
    cancel()
}

func doStuff(ctx context.Context) {
    for {
        time.Sleep(1 * time.Second)
        select {
        case <-ctx.Done():
            logg.Printf("stuff-done")
            return
        default:
            logg.Printf("stuff-work")
        }
    }
}

func doRequest(ctx context.Context) {
    for {
        time.Sleep(time.Second * 1)
        select {
            case <-ctx.Done():
                fmt.Println("request-done")
                return

            default:
                fmt.Println("do-request")
        }
    }
}

func main() {
    logg = log.New(os.Stdout, "", log.Ltime)
    someHandler()
    logg.Printf("down")
}
