package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
)

func main() {
    Get("http://www.baidu.com")
}

func Get(url string) error{
    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return err
    }
    defer response.Body.Close() 
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(body))
    return nil
}
