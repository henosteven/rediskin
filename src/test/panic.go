package main

import (
	"fmt"
	"log"
)

//2017/11/25 13:24:09 [21 22]
//bye~bye~
//come from recover
//recover message: [21 22]
//2017/11/26 08:58:48 panic in othergo

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("come from recover")
			fmt.Println("recover message:", e)
		}
	}()

	defer func() {
		fmt.Println("bye~bye~")
	}()

	go othergo()

	arr := []int{21, 22}
	log.Panic(arr) //先打印信息，然后调用panic(arr)
}

func othergo() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	panic("panic in othergo")
}
