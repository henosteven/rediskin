package main

import (
    "fmt"
    "reflect"
)

type T struct {}

func (t *T) DoSomething() {
    fmt.Println("hello~reflect~struct~method~")
}

func(t T) DoSomethingElse() {
    fmt.Println("just do it~")
}

func(t T) Add(i , j int) int{
    return i + j
}

func main() {
    var t T

    //hello~reflect~struct~method~
    reflect.ValueOf(&t).MethodByName("DoSomething").Call([]reflect.Value{})

    //just do it~
    reflect.ValueOf(t).MethodByName("DoSomethingElse").Call([]reflect.Value{})
    reflect.ValueOf(&t).MethodByName("DoSomethingElse").Call([]reflect.Value{})
    
    var i , j int = 1, 2
    result := reflect.ValueOf(t).MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j)})
    fmt.Println(result) //[<int Value>] 是一个slice
    fmt.Println(result[0].Int()) //3 
}

