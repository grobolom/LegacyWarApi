package main

import (
    "net/http"

    // Third Party
    "gopkg.in/mgo.v2"
    "github.com/julienschmidt/httprouter"

    "./handlers"
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

    r.GET("/leaderboards", handlers.ReadLeaderboards(db))
    r.POST("/leaderboards", handlers.CreateLeaderboard(db))

    r.GET("/tracked_repositories", handlers.ReadTrackedRepositories(db))
    r.POST("/tracked_repositories", handlers.CreateTrackedRepository(db))

    r.GET("/repository_score", handlers.ReadRepositoryScores(db))
    r.POST("/repository_score", handlers.CreateRepositoryScore(db))

    //
    // END OF ENDPOINTS
    //

    http.ListenAndServe(":3001", r)
}
