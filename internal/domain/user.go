package domain


type CreateUserRequest struct {
    ID        string       `json:"id"`
    Name      string    `json:"name"`
    Icon      string    `json:"icon"`
    Age       int       `json:"age"`
    Gender    string   `json:gender`

}


type UpdateUserRequest struct {
    ID        string       `json:"id"`
    Name      string    `json:"name"`
    Icon      string    `json:"icon"`
    Age       int       `json:"age"`
    Gender    string   `json:gender`
}