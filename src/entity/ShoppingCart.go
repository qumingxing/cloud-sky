package entity

import "gopkg.in/mgo.v2/bson"

type ShoppingCart struct {
	PId          bson.ObjectId `bson:"_id"`
	Id          string `json:"id" bson:"id"`
	ProductId          string `json:"productId" bson:"productId"`
	Qty int `json:"qty" bson:"qty"`
	UserToken string `json:"userToken" bson:"userToken"`
}