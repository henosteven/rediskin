package main

import (
    "fmt"
)

func main() {
    const (
        a = iota
        b
        c = "heno"
        d //会延续上一个的操作oo
        e = iota
        f
    )
    //0 1 heno heno 4 5
    fmt.Println(a, b, c, d, e, f)
}
