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
