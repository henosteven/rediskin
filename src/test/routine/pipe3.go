package main

/*
 * 版本 
 * 一个生产方，两个消费方，消费方输出数据merge起来
 * merge出来读取的时候，只读取一个数据，这时候需要通知上游停止写入
 */

import (
    "fmt"
    "sync"
)

func gen(done <-chan struct{}, s ...int)  <-chan int {
    var out = make(chan int)
    go func() { 
        defer close(out)
        for _, n := range s {
            select {
                case out <- n:
                case <-done:
                    return
            }
        }
    } () 
    return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int{
    var out = make(chan int) 
    go func() {
        defer close(out)
        for n := range in {
            select {
                case out <- n * n:
                case <-done:
                    return
            }
        }
    }()
    return out
}

func main() {
    defer close(done)
    done := make(chan struct{})
    in := gen(done, 2, 3, 4, 5)
    c1 := sq(done, in)
    c2 := sq(done, in)
    
    out := merge(done, c1, c2) 
    fmt.Println(<-out)
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    var out = make(chan int)
    var wg sync.WaitGroup 
    output := func(c <-chan int) {
         defer wg.Done()
         for i := range c {
            select {
                case out <- i:
                case <-done:
                    return
            }
         }
         //只读的chan不需要也不允许关闭
         //cannot close receive-only channel
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
