package models

//easyjson:json
type Price struct {
	ID        int64 `json:"id"`
	Value     int   `json:"value"`
	HaircutID int64 `json:"haircut_id"`
}
