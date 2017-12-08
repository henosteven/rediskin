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

type User interface {

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

    var strmap map[string]string
    //output: strmap nil
    if strmap == nil {
        fmt.Println("strmap nil")
    }
    
    strmapref := reflect.ValueOf(strmap)
    //output: strmap nil
    if strmapref.IsNil() {
        fmt.Println("strmap nil")
    }
    
    /*
    IsNil reports whether its argument v is nil. The argument must be a chan, func, interface, map, pointer, or slice value; if it is not, IsNil panics. Note that IsNil is not always equivalent to a regular comparison with nil in Go. For example, if v was created by calling ValueOf with an uninitialized interface variable i, i==nil will be true but v.IsNil will panic as v will be the zero Value.
    要求reflect.ValueOf的参数必须为chan func interface map pointer slice, 如果不是的就会panic
    如果参数是一个为初始化的接口类型变量 i==nil为true, 但是v.IsNil, 因为reflect.ValueOf(v) 就不是上述类型了
    */

    var u User
    //panic: reflect: call of reflect.Value.IsNil on zero Value
    //u就是一个未初始化的一个接口类型变量
    /*
    if reflect.ValueOf(u).IsNil() {
    }
    */

    strmap = make(map[string]string)
    strmapref = reflect.ValueOf(strmap)
    //output: strmap not nil
    if strmapref.IsNil() {
        fmt.Println("strmap nil")
    } else {
        fmt.Println("strmap not nil")
    }

    refm := reflect.ValueOf(t).MethodByName("heno")
    fmt.Println(refm)
    //panic: reflect: call of reflect.Value.IsNil on zero Value
    /*
    if refm.IsNil() {
        fmt.Println("result is nil")
    }
    */
    //output: refm is invalid
    //IsValid reports whether v represents a value. It returns false if v is the zero Value. If IsValid returns false
    if refm.IsValid() {
        fmt.Println("refm is valid")
    } else {
        fmt.Println("refm is invalid")
    }
}
