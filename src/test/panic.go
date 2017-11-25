package main

import (
    "log"
    "fmt"
)

//2017/11/25 13:24:09 [21 22]
//bye~bye~
//come from recover
//recover message: [21 22]

func main() {
    defer func() {
        if e := recover(); e != nil {
            fmt.Println("come from recover")
            fmt.Println("recover message:", e)
        }
    } ()

    defer func() {
        fmt.Println("bye~bye~")
    } ()

    arr := []int{21,22}
    log.Panic(arr) //先打印信息，然后调用panic(arr)
}


