package album

import "time"

type DTO struct {
	Name     string   `json:"name"`
	Genres   []string `json:"genres"`
	Artist   string   `json:"artist""`
	Songs    []string `json:"songs"`
	CreateAt time.Time
	UpdateAt time.Time
}
