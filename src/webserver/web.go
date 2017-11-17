package webserver

import (
    "fmt"
    "net/http"
    "server"
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
    for k, v := range server.ServerInstance.Dict {
        fmt.Fprintf(w, k)
        switch (v.Value).(type) {
            case string:
                fmt.Fprintf(w, v.Value.(string))
                break
            default:
                fmt.Fprintf(w, "not-supprt-v")
                break
        }
    }
    fmt.Fprintf(w, "hello server~")
}
