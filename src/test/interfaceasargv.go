package main

import (
	"fmt"
)

type Listener interface {
	Accept()
	ShowListenedList()
}

type HenoListener struct {
	Listener
}

func (h HenoListener) ShowListenedList() {
	fmt.Println("just show HenoListener")
}

func Job(l Listener) {
	l.ShowListenedList()
}

func main() {
	hl := HenoListener{}
	Job(hl)
}
