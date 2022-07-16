package song

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Song struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Length   time.Duration      `bson:"length"`
	Genres   []string           `bson:"genres"`
	Album    primitive.ObjectID `bson:"album"`
	Artist   primitive.ObjectID `bson:"artist"`
	CreateAt time.Time          `bson:"createAt"`
	UpdateAt time.Time          `bson:"updateAt"`
}

func NewSong(songDTO SongDTO) (*Song, error) {
	var err error
	albumId, err := primitive.ObjectIDFromHex(songDTO.Album)
	artistId, err := primitive.ObjectIDFromHex(songDTO.Artist)
	length, err := time.ParseDuration(songDTO.Length)
	if err != nil {
		return nil, err
	}
	return &Song{
		Name:     songDTO.Name,
		Length:   length,
		Genres:   songDTO.Genres,
		Album:    albumId,
		Artist:   artistId,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}, err
}
