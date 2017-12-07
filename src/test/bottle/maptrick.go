package main

import (
    "fmt"
)

type Handler func()

func main() {
    var handleMap = make(map[string]Handler)
    handleMap["do"] = Do
    handleMap["do"]() //just do something
}

func Do() {
    fmt.Println("just do something")
}
