package models

import "time"

type (
    Leaderboard struct {
        Id              int64               `json:"id"`
        // the start date of the leaderboard
        Start           time.Time           `json:"start"`
        // some leaderboards can be month-long?
        Length          int                 `json:"length"`
        // we can have different ways of scoring these
        ScoringMethod   string              `json:"scoring_method"`
        Board           []RepositoryScore   `json:"board"`
    }

    Leaderboards []Leaderboard
)
