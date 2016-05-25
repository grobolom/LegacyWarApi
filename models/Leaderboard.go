package models

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
