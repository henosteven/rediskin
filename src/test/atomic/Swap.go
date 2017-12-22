package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var a int32 = 32
    var b int32 = 64
    result := atomic.SwapInt32(&a, b)
    fmt.Println(result, a, b)
}
