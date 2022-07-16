package song

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Song struct {
	Id       primitive.ObjectID  `bson:"_id"`
	Name     string              `bson:"name"`
	Length   string              `bson:"length"`
	Genres   []string            `bson:"genres"`
	Album    *primitive.ObjectID `bson:"album,omitempty"`
	Artist   *primitive.ObjectID `bson:"artist"`
	CreateAt time.Time           `bson:"createAt"`
	UpdateAt time.Time           `bson:"updateAt"`
}

func NewSong(songDTO DTO) (*Song, error) {

	albumId, _ := primitive.ObjectIDFromHex(songDTO.Album)
	artistId, err := primitive.ObjectIDFromHex(songDTO.Artist)

	if err != nil {
		return nil, err
	}
	return &Song{
		Id:       primitive.NewObjectID(),
		Name:     songDTO.Name,
		Length:   songDTO.Length,
		Genres:   songDTO.Genres,
		Album:    &albumId,
		Artist:   &artistId,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}, nil
}
