package domain

import "time"

type Caffeine struct {
    ID        string  `json:"id"`
    Score     int  `json:"score"`
    UserID    string  `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}


type CreateCaffeineRequest struct {
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
}

type UpdateCaffeineRequest struct {
    ID        string `json:"id"`
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
}