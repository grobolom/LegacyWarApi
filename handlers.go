package main

import (
    "time"
    "net/http"
    "gopkg.in/mgo.v2"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    "./models"
)


func ReadLeaderboards(db *mgo.Database) httprouter.Handle {
    return func(
        w http.ResponseWriter,
        r *http.Request,
        ps httprouter.Params,
    ) {
        res := []models.Leaderboard{}
        e := db.C("leaderboards").Find(nil).All(&res)
        if e != nil {
            panic(e)
        }
        json.NewEncoder(w).Encode(res)
    }
}

func CreateLeaderboards(db *mgo.Database) httprouter.Handle {
    return func(
        w http.ResponseWriter,
        r *http.Request,
        ps httprouter.Params,
    ) {
        db.C("leaderboards").Insert(&models.Leaderboard{
            Id: "",
            Start: time.Now(),
            Length: 7,
            ScoringMethod: "default",
        })
    }
}
