package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int8
}

func main() {
	u_1 := new(User) //返回指针
	u_1.Name = "heno"
	u_1.Age = 28

	u_2 := User{} //返回值
	u_2.Name = "heno"
	u_2.Age = 28

	//&{heno 28} {heno 28}
	fmt.Println(u_1, u_2)
}
