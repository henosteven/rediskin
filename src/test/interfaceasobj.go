package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() interface{}
	Display()
}

type Rect struct {
	X int
	Y int
}

func (rect *Rect) Area() interface{} {
	return rect.X * rect.Y
}

func (rect *Rect) Display() {
	fmt.Println("x:", rect.X, "y:", rect.Y)
}

type Round struct {
	R int
}

func (round *Round) Area() interface{} {
	return round.R * round.R * int(math.Pi)
}

func (round *Round) Display() {
	fmt.Println("R:", round.R)
}

func Show(shape Shape) {
	fmt.Println("area", shape.Area())
	shape.Display()
}

func main() {
	//Rect does not implement Shape (Area method has pointer receiver)
	//这里主要说明 *SomeType 实现了接口，并不代表 SomeType也实现了接口
	//所以需要使用接口类型的话，就需要使用&VarSomeType
	//rect := Rect{3, 5}

	rect := &Rect{3, 5}
	Show(rect)
	round := &Round{3}
	Show(round)
}
