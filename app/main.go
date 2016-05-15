package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "log"
    "labix.org/v2/mgo"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    session, err := mgo.Dial("legacywarapi_database_1")
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    var result [] struct { name string }
    e := session.DB("local").C("random").Find(nil).All(&result)
    if e != nil {
        panic(e)
    }
    s, e := json.Marshal(result)

    fmt.Fprintf(w, "Go web app + Mongo + Docker!")
    fmt.Println(w, string(s))
}

func main() {
    http.HandleFunc("/", defaultHandler)
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
        return
    }
}
