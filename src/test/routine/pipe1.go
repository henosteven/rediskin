package main

/*
 * 版本 
 */

import (
    "fmt"
)

func gen(s ...int)  <-chan int {
    var out = make(chan int)
    go func() { 
        for _, n := range s {
            out <- n
        }
        close(out) //处理完毕之后关闭写入
    } () 
    return out
}

func sq(in <-chan int) <-chan int{
    var out = make(chan int) 
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out) //处理完毕之后关闭写入
    }()
    return out
}

func main() {
    c := gen(2, 3, 4, 5)
    out := sq(c)
    /*
    for i := range out {
        fmt.Println(i)
    }
    */
    for j := range sq(sq(out)) {
        fmt.Println(j)
    }
}
