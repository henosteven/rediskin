package main

import (
    "sync/atomic"
    "fmt"
)

func main() {
    var a uint32
    atomic.AddUint32(&a, 2)
    fmt.Println(a)
}
