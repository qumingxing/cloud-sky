package entity

import "gopkg.in/mgo.v2/bson"

type Product struct {
	PId          bson.ObjectId `bson:"_id"`
	Id          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Image       string `json:"image" bson:"image"`
	RetailPrice float64 `json:"retailPrice" bson:"retailPrice"`
	Qty int `json:"qty" bson:"qty"`
	Freight float64 `json:"freight" bson:"freight"`
	Descr        string `json:"descr" bson:"descr"`

}