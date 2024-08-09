package domain

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Icon      string    `json:"icon"`
    Age       int       `json:"age"`      
    Gender    string    `json:"gender"`   
}
