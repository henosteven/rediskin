package main

import (
    "fmt"
    "reflect"
    "sync"
)

func main() {

    var once sync.Once
    once.Do(oncefoo)
    once.Do(oncefoo)

    var a uint8 = 8
    c := a * 32 //go不会根据计算之后的类型给c，而是依赖a的类型，这导致存不下
    fmt.Println(c) //这里变成了0
    fmt.Println(reflect.TypeOf(c))//uint8

    slice := make([]interface{}, 10)
    var user map[string]string
    user = make(map[string]string)
    user["name"] = "jinjing"
    slice[0] = user
    slice[1] = "henosteven"
    for _, v := range slice {
        switch t := v.(type) {
            case string:
                fmt.Println("string")
                break
            case map[string]string:
                fmt.Println("map[string]string")
                break
            default:
                fmt.Println("not-support", t)
                break
        }
    }
    echo([]int{1,2,3,4,5})
}

//参数传进来了，但是还是interface{}, 不会做自动转化
//这里需要一个手动转化, 不然会报一个类型不对的错误
func echo(v interface{}) {
    fmt.Println(v)
    cv, ok := v.([]int)
    if !ok {
        fmt.Println("invalid-type")
    }
    for _, tmpv := range cv {
        fmt.Println(tmpv)
    }
}

func oncefoo() {
    fmt.Println("call oncefoo~")
}
