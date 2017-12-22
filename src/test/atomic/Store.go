package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var a int32
    atomic.StoreInt32(&a, 32)
    fmt.Println(a)
}
