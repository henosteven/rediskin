package main

import (
    "fmt"
)

func main() {
    var a uint = 0
    a--

    //yes we get this: 18446744073709551615
    //so len(slice) return int type instead of uint, though the result will never be negative
    //because we may use where len this way: i := len(slice); i >= 0 ; i--
    //then the i will out of range, which will lead to a panic
    fmt.Println(a) 
}
