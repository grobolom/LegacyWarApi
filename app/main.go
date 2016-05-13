package main

import (
    "net/http"
    "fmt"
    "log"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Go web app + Mongo + Docker!")
}

func main() {
    http.HandleFunc("/", defaultHandler)
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
        return
    }
}
