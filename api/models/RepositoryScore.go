package models

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
