package main

import (
    "net/http"
    "fmt"
    "strconv"
)

type database map[string]int

func (d *database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    for k , v := range *d {
        fmt.Fprintf(w, k + " : " + strconv.Itoa(v))
    }
}

func main() {
    var data  = &database {
        "tool" : 101,
        "sock" : 100,
    } 
    
    server := &http.Server{Addr:":8080", Handler: data}
    server.ListenAndServe()
}
