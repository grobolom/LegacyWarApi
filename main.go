package main

import (
    "encoding/json"
    "net/http"

    // Third Party
    "gopkg.in/mgo.v2"
    "github.com/julienschmidt/httprouter"

)

import "./models"

func main() {
    mgo_url := ":27017"
    session, err := mgo.Dial(mgo_url)
    c := session.DB("legacywar-api").C("leaderboards")
    _ = c
    _ = err

    // Instantiate router
    r := httprouter.New()

    // set up a GET handler
    r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        res := []models.Leaderboard{}
        err := c.Find(nil).All(&res)
        if err != nil {}
        json.NewEncoder(w).Encode(res)
    })

    // httprouter does cool stuff I think
    http.ListenAndServe(":3000", r)
}
