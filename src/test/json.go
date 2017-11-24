package main

import (
    "fmt"
    "encoding/json"
)

type User struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func main() {
    user := User{Name:"jinjing", Age:28}
    b, _ := json.Marshal(user)
    fmt.Println(string(b)) //{"name":"jinjing","age":28}

    var u User
    json.Unmarshal(b, &u)
    fmt.Println(u) //{jinjing 28}

    var u2 User
    json.Unmarshal([]byte("{\"name\":\"henosteven\",\"age\":28}"), &u2)
    fmt.Println(u2) //{henosteven 28}
}
