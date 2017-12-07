package main

/*
 * 版本 
 * 一个生产方，两个消费方，消费方输出数据merge起来
 */

import (
    "fmt"
    "sync"
)

func gen(s ...int)  <-chan int {
    var out = make(chan int)
    go func() { 
        for _, n := range s {
            out <- n
        }
        close(out)
    } () 
    return out
}

func sq(in <-chan int) <-chan int{
    var out = make(chan int) 
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    in := gen(2, 3, 4, 5)
    c1 := sq(in)
    c2 := sq(in)

    for i := range merge(c1, c2) {
        fmt.Println(i)
    }
}

func merge(cs ...<-chan int) <-chan int {
    var out = make(chan int)
    var wg sync.WaitGroup 
    output := func(c <-chan int) {
         for i := range c {
            out <- i
         }
         //只读的chan不需要也不允许关闭
         //cannot close receive-only channel
         wg.Done()
    }
    
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }
    
    go func() {
        wg.Wait()
        close(out) //处理完毕之后关闭写入
    } ()
    
    return out
}
