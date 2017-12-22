package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var v uint32 = 32
    tmpv := atomic.LoadUint32(&v)
    fmt.Println(tmpv)
}
