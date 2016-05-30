package entity

import "gopkg.in/mgo.v2/bson"
/**
广告
 */
type Adverts struct {
	PId          bson.ObjectId `bson:"_id"`
	Id          string `json:"id" bson:"id"`
	Image string `json:"image" bson:"image"`
	Title string `json:"title" bson:"title"`
	ProductId string `json:"productId" bson:"productId"`
}
