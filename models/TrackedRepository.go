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
