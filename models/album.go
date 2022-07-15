package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Album struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Artist Artist        `json:"artist" bson:"artist""`
	Songs  []Song        `json:"songs" bson:"songs"`
}
