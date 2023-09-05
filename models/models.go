package models

import (
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	//Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name          string             `json:"name,omitempty" bson:"name"`
	Phoneno       int                `json:"no" bson:"no"`
	Profilestatus bool               `json:"status" bson:"ststus"`
	Transaction   []Transaction      `json:"transaction" bson:"transaction"`
	Balance       float64            `json:"balance" bson:"balance"`
}
type Transaction struct {
	Type           string    `json:"type" bson:"type"`
	From_coustomer string    `json:"from" bson:"from"`
	To_coustomer   string    `json:"to" bson:"to"`
	Amount         float64   `json:"amount" bson:"amount"`
	Timestamp      time.Time `jsom:"time" bson:"time"`
}

/*
type Profile struct {
	Coustomername   string    `json:"name" bson:"name"`
	TransactionDate time.Time `json:"date" bson:"date"`
	TransactionAmount int `json:"amount" bson:"amount"`

}*/
