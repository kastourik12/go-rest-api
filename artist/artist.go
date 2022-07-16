package artist

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Artist struct {
	Id       primitive.ObjectID   `bson:"_id"`
	Name     string               `bson:"name"`
	Age      string               `bson:"age"`
	ImageURL string               `bson:"imageURL"`
	Genres   []string             `bson:"genres"`
	Albums   []primitive.ObjectID `bson:"albums"`
	Songs    []primitive.ObjectID `bson:"songs"`
	CreateAt time.Time            `bson:"createAt"`
	UpdateAt time.Time            `bson:"updateAt"`
}
