package song

import (
	"time"
)

type SongDTO struct {
	Name     string        `json:"name"`
	Length   time.Duration `json:"length"`
	Genres   []string      `json:"genres"`
	Album    string        `json:"album"`
	Artist   string        `json:"artist"`
	CreateAt time.Time
	UpdateAt time.Time
}
