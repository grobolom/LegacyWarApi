package handlers

import (
    "net/http"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    "../models"
)


func ReadTrackedRepositories(db *mgo.Database) httprouter.Handle {
    return func(
        w http.ResponseWriter,
        r *http.Request,
        ps httprouter.Params,
    ) {
        res := []models.TrackedRepository{}
        e := db.C("tracked_repositories").Find(nil).All(&res)
        if e != nil {
            panic(e)
        }
        json.NewEncoder(w).Encode(res)
    }
}

func CreateTrackedRepository(db *mgo.Database) httprouter.Handle {
    return func(
        w http.ResponseWriter,
        r *http.Request,
        ps httprouter.Params,
    ) {
        n := models.TrackedRepository{}
        err := json.NewDecoder(r.Body).Decode(&n)

        if err != nil {
            panic(err)
        }

        n.Id = bson.NewObjectId()

        db.C("tracked_repositories").Insert(&n)
        json.NewEncoder(w).Encode(n)
    }
}
