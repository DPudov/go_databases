package models

//easyjson:json
type Employee struct {
	ID        int64  `json:"id"`
	FullName  string `json:"fullname"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IsWorking bool   `json:"is_working"`
	SalonID   int64  `json:"salon_id"`
}
