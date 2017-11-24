package main

import (
    "fmt"
    "regexp"
    "reflect"
)

func main() {
    pat := `h(.*?)o`
    re := regexp.MustCompile(pat)
    fmt.Println(reflect.TypeOf(re)) //*regexp.Regexp
    result := re.Find([]byte("hello heno steven"))
    fmt.Println(string(result)) //hello
    fmt.Println(result) //[104 101 108 108 111]
    fmt.Println(reflect.TypeOf(result)) //[]uint8

    allResult := re.FindAll([]byte("hello heno steven"), -1)
    fmt.Println(allResult) //[[104 101 108 108 111] [104 101 110 111]]
    fmt.Println(reflect.TypeOf(allResult)) //[][]uint8

    fmt.Println(re.FindAllString("hello heno steven", -1)) //[hello heno]
    fmt.Println(re.FindStringIndex("i~hello heno steven")) //[2 7]
    fmt.Println(re.FindAllStringIndex("i~hello heno steven", -1)) //[[2 7] [8 12]]
    fmt.Println(re.FindIndex([]byte("i~hello heno steven"))) //[2 7]
    fmt.Println(re.FindAllIndex([]byte("i~hello heno steven"), -1)) //[[2 7] [8 12]]

    fmt.Println(re.FindSubmatch([]byte("i~hello heno steven"))) //[[104 101 108 108 111] [101 108 108]]
    fmt.Println(re.FindStringSubmatch("i~hello heno steven")) //[hello ell]
}
