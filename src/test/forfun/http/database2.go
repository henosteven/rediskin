package main

import (
    "net/http"
    "fmt"
    "strconv"
)

type database map[string]int

func (d *database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
        case "/" :
            for k , v := range *d {
                fmt.Fprintf(w, k + " : " + strconv.Itoa(v))
            }
        case "/price":
            item := r.URL.Query().Get("item")
            price, ok := (*d)[item]
            if !ok {
                w.WriteHeader(http.StatusNotFound)
                fmt.Fprintf(w, "item not exist");
            } else {
                fmt.Fprintf(w, item + ":" + strconv.Itoa(price))
            }
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
