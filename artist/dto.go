package artist

import "time"

type DTO struct {
	Name     string   `json:"name"`
	Age      string   `json:"age"`
	ImageURL string   `json:"imageURL"`
	Genres   []string `json:"genres"`
	CreateAt time.Time
	UpdateAt time.Time
}
