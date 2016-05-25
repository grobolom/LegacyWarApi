package models

import "time"

type (
    TrackedRepository struct {
        id int64
        name string
        owner string
        full_name string
        url string
        included_since time.Time
        active bool
        active_since time.Time
    }
)
