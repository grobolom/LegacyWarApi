package handlers

import (
    "net/http"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    "../models"
)


func ReadRepositoryScores(db *mgo.Database) httprouter.Handle {
    return func(
        w http.ResponseWriter,
        r *http.Request,
        ps httprouter.Params,
    ) {
        res := []models.RepositoryScore{}
        e := db.C("repository_scores").Find(nil).All(&res)
        if e != nil {
            panic(e)
        }
        json.NewEncoder(w).Encode(res)
    }
}

func CreateRepositoryScore(db *mgo.Database) httprouter.Handle {
    return func(
        w http.ResponseWriter,
        r *http.Request,
        ps httprouter.Params,
    ) {
        n := models.RepositoryScore{}
        err := json.NewDecoder(r.Body).Decode(&n)

        if err != nil {
            panic(err)
        }

        n.Id = bson.NewObjectId()

        db.C("repository_scores").Insert(&n)
        json.NewEncoder(w).Encode(n)
    }
}
