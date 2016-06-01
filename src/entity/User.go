package entity

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	PId          bson.ObjectId `bson:"_id"`
	UserId          string `json:"userId" bson:"userId"`
	Mobile     string `json:"mobile" bson:"mobile"`
	Password   string `json:"password" bson:"password"`
	UserToken  string `json:"userToken" bson:"userToken"`
	Status     string `json:"status" bson:"status"`
	RealName   string `json:"realName" bson:"realName"`
	Image      string `json:"image" bson:"image"`
	CreateDate time.Time `json:"createDate" bson:"createDate"`
}
