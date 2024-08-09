package domain

import "time"

type Sleep struct {
    ID        string    `json:"id"`
    Score     int    `json:"score"`
    UserID    string    `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}



type CreateSleepRequest struct {
    Score     string `json:"score"`
    UserID    string `json:"user_id"`
}

type UpdateSleepRequest struct {
    ID        string `json:"id"`
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
}