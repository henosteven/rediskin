package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string            `json:"name"`
	Age  int               `json:"age"`
	Ext  map[string]string `json:"userext"`
}

func main() {

	testm := make(map[string]string)
	testm["name"] = "jinjing"
	testm["avatar"] = "http"
	tmpJson, _ := json.Marshal(testm)
	fmt.Println(string(tmpJson)) //{"avatar":"http","name":"jinjing"}

	tmpJson, _ = json.Marshal([]int{1, 2, 3, 4, 5})
	fmt.Println(string(tmpJson)) //[1,2,3,4,5]

	user := User{Name: "jinjing", Age: 28}
	user.Ext = testm
	b, _ := json.Marshal(user)
	fmt.Println(string(b)) //{"name":"jinjing","age":28,"userext":{"avatar":"http","name":"jinjing"}}

	var u User
	json.Unmarshal(b, &u)
	fmt.Println(u) //{jinjing 28}

	var u2 User
	json.Unmarshal([]byte("{\"name\":\"henosteven\",\"age\":28}"), &u2)
	fmt.Println(u2) //{henosteven 28}
}
