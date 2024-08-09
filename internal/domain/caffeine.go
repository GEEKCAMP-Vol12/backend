package domain



type Caffeine struct {
    ID        string `json:"id"`
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
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