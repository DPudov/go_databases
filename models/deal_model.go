package models

import "time"

//easyjson:json
type Deal struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customer_id"`
	HaircutID  int64     `json:"haircut_id"`
	EmployeeID int64     `json:"employee_id"`
	PriceID    int64     `json:"price_id"`
	Date       time.Time `json:"date"`
}
