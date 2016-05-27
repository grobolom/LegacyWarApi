package models

import "gopkg.in/mgo.v2/bson"

type (
    RepositoryScore struct {
        Id bson.ObjectId `json:"_id" bson:"_id"`
        Name string `json:"name" bson:"name"`
        Owner string `json:"owner" bson:"owner"`
        FullName string `json:"fullname" bson:"fullname"`
        ScoringMethod string `json:"scoring_method" bson:"scoring_method"`
        Score int `json:"score" bson:"score"`
        Position int `json:"position" bson:"position"`
    }

    RepositoryScores []RepositoryScore
)
