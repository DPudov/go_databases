package models

//easyjson:json
type Customer struct {
	ID       int64  `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
}
