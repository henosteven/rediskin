package main

import (
	"fmt"
	"sync"
	"time"
)

type safeInt struct {
	sync.Mutex
	Num int
}

var a safeInt
var ch chan bool

func main() {
	a = safeInt{Num: 5}
	ch = make(chan bool)
	go test()
	go func() {
		for i := 0; i < 10; i++ {
			a.Lock()
			sum := i + a.Num
			fmt.Println("before-sleep", a.Num)
			time.Sleep(time.Second)
			fmt.Println(i, a.Num, sum)
			fmt.Println("========")
			a.Unlock()
			ch <- true
		}
	}()
	for i := 0; i < 20; i++ {
		<-ch
	}
}

func test() {
	for i := 0; i < 10; i++ {
		a.Num = i
		time.Sleep(time.Second)
		ch <- true
	}
}
