package main

import (
	"fmt"
)

type User struct {
	Name string
}

func (u *User) showName() {
	fmt.Println(u.Name)
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u User) getName() string {
	return u.Name
}

func main() {
	var u1 = User{}
	u1.setName("jinjing")
	u1.showName()    //这里会自动变成(&u1).showName
	(&u1).showName() //等价

	u2 := &u1
	fmt.Println(u2.getName())    //这里会自动转化为(*u2).getName
	fmt.Println((*u2).getName()) //等价
}
