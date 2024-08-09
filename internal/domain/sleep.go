package domain


type Sleep struct {
    ID        string `json:"id"`
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    
}


type CreateSleepRequest struct {
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
}

type UpdateSleepRequest struct {
    ID        string `json:"id"`
    Score     int `json:"score"`
    UserID    string `json:"user_id"`
}