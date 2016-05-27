package models

import "time"
import "gopkg.in/mgo.v2/bson"

type (
    TrackedRepository struct {
        Id bson.ObjectId `json:"_id" bson:"_id"`
        Name string `json:"name" bson:"name"`
        Owner string `json:"owner" bson:"owner"`
        FullName string `json:"fullname" bson:"fullname"`
        Url string `json:"url" bson:"url"`
        IncludedSince time.Time `json:"included_since" bson:"included_since"`
        Active bool `json:"active" bson:"active"`
        ActiveSince time.Time `json:"active_since" bson:"active_since"`
    }

    TrackedRepositories []TrackedRepository
)
