package domain

import "time"

type Sleep struct {
    ID        int    `json:"id"` 
    Score     int    `json:"score"`
    UserID    int     `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

