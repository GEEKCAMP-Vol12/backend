package domain




type CreateCaffeineRequest struct {
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
}

type UpdateCaffeineRequest struct {
    ID        string `json:"id"`
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
}