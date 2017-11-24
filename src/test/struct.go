package main

import (
    "fmt"
)

type User struct {
    Name string
}

//全局变量
var user User

func main() {
    user = User{Name: "jinjing"}
    fmt.Println(user)
    setUserName("henosteven")
    fmt.Println(user)
}

func setUserName(name string) {
    //全局变量可以被用到
    user.Name = name
}
