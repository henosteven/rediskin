package main

import (
    "fmt"
    "regexp"
    "reflect"
)

func main() {
    pat := `h.*?o`
    re := regexp.MustCompile(pat)
    fmt.Println(reflect.TypeOf(re)) //*regexp.Regexp
    result := re.Find([]byte("hello heno steven"))
    fmt.Println(string(result)) //hello
    fmt.Println(result) //[104 101 108 108 111]
    fmt.Println(reflect.TypeOf(result)) //[]uint8

    allResult := re.FindAll([]byte("hello heno steven"), -1)
    fmt.Println(allResult) //[[104 101 108 108 111] [104 101 110 111]]
    fmt.Println(reflect.TypeOf(allResult)) //[][]uint8
}
