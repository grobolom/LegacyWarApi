package main

import (
    "net/http"

    // Third Party
    "gopkg.in/mgo.v2"
    "github.com/julienschmidt/httprouter"

)

func main() {
    // MGO SETUP
    mgo_url := ":27017"
    session, err := mgo.Dial(mgo_url)
    if err != nil {
        panic(err)
    }
    db := session.DB("legacywar-api")

    r := httprouter.New()

    //
    // ENDPOINTS
    //

    r.GET("/leaderboards", ReadLeaderboards(db))
    r.POST("/leaderboards", CreateLeaderboards(db))
    http.ListenAndServe(":3000", r)

    //
    // END OF ENDPOINTS
    //

}
