package main

import (
    "fmt"
)

func appendInt(s []int, y int) []int{
    var z []int
    zlen := len(s) + 1
    if zlen <= cap(s) {
        z = s[:zlen] 
    } else {
        zcap := zlen
        if zcap < 2 * len(s) {
            zcap =  2 * len(s)
        }
        z = make([]int, zlen, zcap)
        copy(z, s)
    }
    z[zlen - 1] = y
    return z
}

func main() {
    s := []int{1, 2, 3}
    s = appendInt(s, 4)
    s = appendInt(s, 5)
    fmt.Println(s)
}
