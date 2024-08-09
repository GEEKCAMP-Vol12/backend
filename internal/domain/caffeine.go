package domain

import "time"

type Caffeine struct {
    ID        int  `autoIncrement:"true" json:"id"`
    Score     int  `json:"score"`
    UserID    int  `json:"user_id"`
    CreatedAt time.Time `time.Time:"true" json:"created_at"`
    UpdatedAt time.Time `time.Time:"true" json:"updated_at"`
}
