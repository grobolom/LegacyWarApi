package main

import (
    "net/http"
    "fmt"
    "log"
    "labix.org/v2/mgo"
    "os"
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

// putting things in here for now while we figure out how to work with
// mongo better

func db() *mgo.Session {
    session, err := mgo.Dial("legacywarapi_database_1")
    if err != nil {
        panic(err)
    }
    session.SetMode(mgo.Monotonic, true)
    session.DB('local')
    return session
}
