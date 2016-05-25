package main

import (
    "fmt"
    "net/http"
    // "encoding/json"
    "labix.org/v2/mgo"

    // Third Party
    "github.com/julienschmidt/httprouter"
)

func defaultHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "Go web app + Mongo + Docker!")
}

func main() {
    // set up Mongo
    session, err := mgo.Dial("legacywarapi_database_1")
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    // Instantiate router
    r := httprouter.New()

    // set up a GET handler
    r.GET("/", defaultHandler)

    // httprouter does cool stuff I think
    http.ListenAndServe("localhost:3000", r)
}
