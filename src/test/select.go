package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- "henosteven"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		ch2 <- "jing"
	}()

	//Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.
	select {
	case res := <-ch1:
		fmt.Println("recv from ch", res)
	case res := <-ch2:
		fmt.Println("recv from ch", res)
	}
}
