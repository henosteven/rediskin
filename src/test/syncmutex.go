package main

import (
    "fmt"
    "time"
    "sync"
)

type safeInt struct {
    sync.Mutex
    Num int
}

var a  safeInt
var wg sync.WaitGroup
func main() {
    a  = safeInt{Num: 5}
    wg.Add(1)
    go test()
    go func() {
        wg.Add(1)
        defer wg.Done()
        for i := 0; i < 10; i++ {
            a.Lock()
            sum := i + a.Num
            fmt.Println("before-sleep", a.Num)
            time.Sleep(time.Second)
            fmt.Println(i, a.Num , sum)
            fmt.Println("========")
            a.Unlock()
        }
    }()
    wg.Wait()
}

func test() {
    wg.Add(1)
    defer wg.Done()
    for i := 0; i < 10 ; i++ {
        a.Num = i
        time.Sleep(time.Second)
    }
}
