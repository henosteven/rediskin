package main

import (
	"fmt"
)

//定义函数类型
type Handler func(input string) error

func h(input string) error {
	fmt.Println(input)
	return nil
}

//类型都可以作为recive挂方法
func (h Handler) doSomething(input string) {
	h(input)
}

func call(input string, handler Handler) {
	handler.doSomething(input)
}

func main() {
	call("heno", Handler(h))
}
