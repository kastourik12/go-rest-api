package album

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Album struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Genres   []string           `bson:"genres"`
	Artist   primitive.ObjectID `bson:"artist""`
	Songs    primitive.ObjectID `bson:"songs"`
	CreateAt time.Time          `bson:"createAt"`
	UpdateAt time.Time          `bson:"updateAt"`
}
