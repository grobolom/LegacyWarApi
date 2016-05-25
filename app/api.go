/*

    Preliminary Structs for REST API endpoints
*/
package models

type (
    TrackedRepository struct {
        id int64
        name string
        owner string
        full_name string
        url string
        included_since time // import "time"
        active bool
        active_since time
    }
)

type (
    RepositoryScore struct {
        repository_id struct
        name string
        owner string
        full_name string
        scoring_method string
        score int64
        position int
    }
)

type (
    Leaderboard struct {
        id int64
        start time // the start date of the leaderboard
        length int // some leaderboards can be month-long?
        score int64 // maybe float? we can probably round it in our calculations
        scoring_method string // we can have different ways of scoring these
        board []RepositoryScore
    }
)
