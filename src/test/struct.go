package main

import (
    "fmt"
)

/*
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

只能为本包之内的类型申明函数，不能为非本包内的类型包括内建类型申明函数

cannot define new methods on non-local type int
func (i int) test() {
}
*/

type User struct {
    Name string
    Age int
}

func (u *User) DisplayName() {
    fmt.Println(u.Name, u.Age)
}

func (u *User) SetAge(age int) {
    u.Age = age
}

/*
With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.)
跟方法的参数一样，如果是一个value receiver，那么在函数内部就是一份拷贝
*/
func (user User) setUserNameValue(name string) {
    user.Name = name
}

/*
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
*/
func (user *User)setUserNamePointer(name string) {
    user.Name = name
}

func main() {
    user2 := User{}
    setUserNameValue("xiaopang", user2)
    user2.DisplayName() // 0
    setUserNamePointer("xiaopang", &user2)
    user2.DisplayName() // xiaopang 0
    user2.setUserNameValue("jinjing")
    user2.DisplayName() // xiaopang 0
    user2.setUserNamePointer("jinjing")
    user2.DisplayName() //jinjing 0
}

func setUserNameValue(name string, user User) {
    user.Name = name
}

func setUserNamePointer(name string, user *User) {
    user.Name = name
}
