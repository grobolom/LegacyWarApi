package models

import "time"
import "gopkg.in/mgo.v2/bson"

type (
    Leaderboard struct {
        Id bson.ObjectId `json:"_id" bson:"_id"`
        // the start date of the leaderboard
        Start time.Time `json:"start" bson:"start"`
        // some leaderboards can be month-long?
        Length int `json:"length" bson:"length"`
        // we can have different ways of scoring these
        ScoringMethod string `json:"scoring_method" bson:"scoring_method"`
        Board []RepositoryScore `json:"board" bson:"board"`
    }

    Leaderboards []Leaderboard
)
