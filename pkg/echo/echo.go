package echo

import (
    "fmt"
    "net/http"
)

func StartServer() {
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Echo: ", r.URL.Path[1:])
    })
    http.ListenAndServe(":8080", nil)
}