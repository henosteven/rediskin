package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var pipe sync.Pool
	pipe = sync.Pool{
		New: func() interface{} {
			return "hello Beijing"
		},
	}

	ch := make(chan interface{}, 2)
	tmp := "hello world"
	pipe.Put(tmp)
	for i := 0; i < 2; i++ {
		go func() {
			ch <- pipe.Get()
		}()
	}

	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}

	var bytePool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1024)
			return &b
		},
	}
	a := time.Now().Unix()
	for i := 0; i < 1000000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().Unix()
	for i := 0; i < 1000000000; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("without-pool", b-a) //53
	fmt.Println("with-pool", c-b)    //39
}
