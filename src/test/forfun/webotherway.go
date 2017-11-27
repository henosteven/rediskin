package main

import (
    "net/http"
    "fmt"
    "reflect"
    "context"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello server")
    fmt.Println(r.Context())
}

func main() {
    serveMux := http.NewServeMux()
    serveMux.HandleFunc("/", index)
    fmt.Println(reflect.TypeOf(serveMux)) //*http.ServeMux
    
    newServeMux := middleWare(serveMux)
    server := http.Server{Addr: ":8080", Handler:newServeMux}
    server.ListenAndServe()
}

func middleWare(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       pctx := r.Context()
       //cctx := context.WithValue(pctx, []int{1,2}, "henosteven") //key is not comparable panic
       //没有任何意义，就是测试而已,可以用来做个统一的参数处理，例如放入一个traceid
       cctx := context.WithValue(pctx, [2]int{1,2}, "henosteven") //key is not comparable panic
       next.ServeHTTP(w, r.WithContext(cctx)) 
    })
}
