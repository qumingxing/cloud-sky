package entity

type Address struct {
	Name    string `json:"name" bson:"name"`
	Phone   string `json:"phone" bson:"phone"`
	Provice string `json:"provice" bson:"provice"`
	City    string `json:"city" bson:"city"`
	Region  string `json:"region" bson:"region"`
	Street  string `json:"street" bson:"street"`
}
