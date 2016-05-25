package main

import (
    "net/http"

    // Third Party
    "gopkg.in/mgo.v2"
    "github.com/julienschmidt/httprouter"

)

func main() {
    // set up mgo here and pass it in to handler functions
    mgo_url := ":27017"
    session, err := mgo.Dial(mgo_url)
    if err != nil {
        panic(err)
    }
    db := session.DB("legacywar-api")

    r := httprouter.New()
    r.GET("/leaderboards", ReadLeaderboards(db))
    http.ListenAndServe(":3000", r)
}
