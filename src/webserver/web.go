package webserver

import (
    "fmt"
    "net/http"
    "server"
    "strconv"
    "service"
)

func StartWebServer() {
    defer service.Wg.Done()
    http.HandleFunc("/serverinfo", serverinfo)
    err := http.ListenAndServe(":" + strconv.Itoa(server.ServerInstance.Webport), nil)
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
