package song

type SongDTO struct {
	Name   string   `json:"name"`
	Length string   `json:"length"`
	Genres []string `json:"genres"`
	Album  string   `json:"album"`
	Artist string   `json:"artist"`
}
