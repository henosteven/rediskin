package main

import (
    "net/http"
    "fmt"
    "strconv"
)

type database map[string]int

func (d *database) list(w http.ResponseWriter, r *http.Request) {
    for k , v := range *d {
        fmt.Fprintf(w, k + " : " + strconv.Itoa(v))
    }
}

func (d *database) price(w http.ResponseWriter, r *http.Request) {
    item := r.URL.Query().Get("item")
    price, ok := (*d)[item]
    if !ok {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "item not exist");
    } else {
        fmt.Fprintf(w, item + ":" + strconv.Itoa(price))
    }
}

func main() {
    var data  = &database {
        "tool" : 101,
        "sock" : 100,
    } 
    
    servermux := http.NewServeMux()
    servermux.Handle("/", http.HandlerFunc(data.list))
    servermux.Handle("/price", http.HandlerFunc(data.price))
    server := &http.Server{Addr:":8080", Handler: servermux}
    server.ListenAndServe()
}
