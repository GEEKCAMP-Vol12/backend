package domain



type CreateSleepRequest struct {
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
}

type UpdateSleepRequest struct {
    ID        string `json:"id"`
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
}