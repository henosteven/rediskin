package webserver

import (
    "fmt"
    "net/http"
)

func StartWebServer() {
    http.HandleFunc("/", serverinfo)
    err := http.ListenAndServe(":1008", nil)
    if err != nil {
        fmt.Println("failed to start webserver")
    }
}

func serverinfo(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Fprintf(w, "hello server~")
}
