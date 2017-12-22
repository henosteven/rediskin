package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var a uint32 = 32
    result := atomic.CompareAndSwapUint32(&a, 32, 64)
    fmt.Println(a, result)
}
