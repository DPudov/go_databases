package models

import "time"

//easyjson:json
type Price struct {
	ID        int64     `json:"id"`
	Value     int       `json:"value"`
	HaircutID int64     `json:"haircut_id"`
	Date      time.Time `json:"date"`
}
