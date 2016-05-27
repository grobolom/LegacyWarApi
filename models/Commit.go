package models

import "time"
import "gopkg.in/mgo.v2/bson"

type (
    Commit struct {
        Id bson.ObjectId `json:"_id" bson:"_id"`
        Url string `json:"url" bson:"url"`
        Date time.Time `json:"date" bson:"date"`
        Author string `json:"author" bson:"author"`
        RepositoryId bson.ObjectId `json:"repository_id" bson:"repository_id"`
        Message string `json:"message" bson:"message"`
        Additions int `json:"additions" bson:"additions"`
        Deletions int `json:"deletions" bson:"deletions"`
        Changes Changes `json:"changes" bson:"changes"`
        ImprovesCoverage bool `json:"improves_coverage" bson:"improves_coverage"`
        ReducesCode bool `json:"reduces_code" bson:"reduces_code"`
    }

    Commits []Commit

    Change struct {
        Filename string `json:"filename" bson:"filename"`
        Additions string `json:"additions" bson:"additions"`
        Deletions string `json:"deletions" bson:"deletions"`
    }

    Changes []Change
)
