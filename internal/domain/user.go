package domain

import "time"

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Icon      string    `json:"icon"`
    Age       int       `json:"age"`      
    Gender    string    `json:"gender"`   
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
