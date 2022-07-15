package models

import (
	"gopkg.in/mgo.v2/bson"
	_ "gopkg.in/mgo.v2/bson"
)

type Song struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Artist string        `json:"artist" bson:"artist"`
	Length string        `json:"length" bson:"length"`
}
