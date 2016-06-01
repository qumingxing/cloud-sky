package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Order struct {
	PId         bson.ObjectId `bson:"_id"`
	OrderId     string `json:"orderId" bson:"orderId"`
	Status      string `json:"status" bson:"status"`
	//商品金额
	Amount      float64 `json:"amount" bson:"amount"`
	//运费
	Freight     float64 `json:"freight" bson:"freight"`
	//总额
	TotalAmount float64 `json:"totalAmount" bson:"totalAmount"`
	CreateDate  time.Time `json:"createDate" bson:"createDate"`
	UserId      string `json:"userId" bson:"userId"`
	Products    []Product `json:"products" bson:"products"`
	//收货地址
	Address     *Address `json:"address" bson:"address"`
}
