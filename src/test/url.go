package main

import (
    "fmt"
    "net/url"
    "reflect"
)

func main() {
    u, err := url.Parse("http://www.heno.me/search?name=jinjing")
    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println(u)
    u.Scheme = "https"
    u.Host = "google.com"
    q := u.Query()
    q.Set("q", "golang")
    u.RawQuery = q.Encode()
    fmt.Println(u)
    fmt.Println(reflect.TypeOf(u)) //*url.URL
    fmt.Println(reflect.ValueOf(u))
}
