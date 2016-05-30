package entity

import "gopkg.in/mgo.v2/bson"

type Category struct {
	PId          bson.ObjectId `bson:"_id"`
	Id          string `json:"id" bson:"id"`
	Image string `json:"image" bson:"image"`
	Name  string `json:"name" bson:"name"`
}