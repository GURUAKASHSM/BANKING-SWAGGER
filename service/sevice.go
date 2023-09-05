package service

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mongoapi/config"
	"mongoapi/models"
)

func Insert(profile models.Profile) {
	inserted, err := config.Collection.InsertOne(context.Background(), profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted", inserted.InsertedID)
}

type BankAccount struct {
	balance1 float64
	balance2 float64
}

func Deposit(amount float64, objectid1 string, objectid2 string) string {
	var acc BankAccount // Initialize BankAccount outside the function

	id1, err := primitive.ObjectIDFromHex(objectid1)
	if err != nil {
		log.Fatal(err)
	}
	id2, err := primitive.ObjectIDFromHex(objectid2)
	if err != nil {
		log.Fatal(err)
	}
	filter1 := bson.M{"_id": id1}
	var account models.Profile
	err = config.Collection.FindOne(context.Background(), filter1).Decode(&account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Senders Current balance: %.2f\n", account.Balance)
	fmt.Printf("Depositing Amount %.2f\n", amount)
	var str = "Insufficient Funds"
	if acc.balance1 < amount {
		return str
	}
	acc.balance1 = account.Balance - amount // Initialize acc's balance
	fmt.Printf("Senders New balance: %.2f\n", acc.balance1)
	update1 := bson.M{"$set": bson.M{"balance": acc.balance1}}
	config.Collection.UpdateOne(context.Background(), filter1, update1)
	if err != nil {
		log.Fatal(err)
	}
	filter2 := bson.M{"_id": id2}

	err = config.Collection.FindOne(context.Background(), filter2).Decode(&account)
	if err != nil {
		log.Fatal(err)
	}
	acc.balance2 = account.Balance + amount

	fmt.Printf("Receiver Current balance: %.2f\n", account.Balance)
	fmt.Printf("Crediting %.2f\n", amount)
	fmt.Printf("Receiver New balance: %.2f\n", acc.balance2)

	update2 := bson.M{"$set": bson.M{"balance": acc.balance2}}
	config.Collection.UpdateOne(context.Background(), filter2, update2)
	if err != nil {
		log.Fatal(err)
	}
	str = "Transfer Successful"
	return str
}

func UpdateOne(objectid string) {
	id, err := primitive.ObjectIDFromHex(objectid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"profilestatus": true}}
	result, err := config.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result of updated", result.MatchedCount)
}

func DeleteOne(objectid string) {
	id, err := primitive.ObjectIDFromHex(objectid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	delete, err := config.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", delete.DeletedCount)
}

func DeleteAll() int64 {
	filter := bson.D{}
	delete, err := config.Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Many:", delete.DeletedCount)
	return delete.DeletedCount
}

func Getalldatabydate(fromDate, toDate time.Time) []primitive.M {
	filter := bson.M{
		"date": bson.M{
			"$gte": fromDate,
			"$lte": toDate,
		},
	}

	cursor, err := config.Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []primitive.M
	for cursor.Next(context.Background()) {
		var profile bson.M
		err := cursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)
	}

	return Profiles
}

func Getsumbydate(fromDate, toDate time.Time) (interface{}, error) {
	filter := bson.M{
		"date": bson.M{
			"$gte": fromDate,
			"$lte": toDate,
		},
	}

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: filter},
		},
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "totalAmount", Value: bson.D{
					{Key: "$sum", Value: "$amount"},
				}},
				{Key: "data", Value: bson.D{
					{Key: "$push", Value: "$$ROOT"},
				}},
			}},
		},
	}

	cursor, err := config.Collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return 0, nil
	}
	defer cursor.Close(context.Background())

	var result []bson.M
	if err := cursor.All(context.Background(), &result); err != nil {
		return 0, nil
	}

	totalAmount := result[0]["totalAmount"]

	fmt.Println(totalAmount)
	return totalAmount, nil
}
// @Summary Get a list of items
// @Description Get a list of items
// @ID get-list-of-items
// @Produce json
// @Success 200 {object} []Item
// @Router /items [get]
func Getalldata() []primitive.M {
	filter := bson.D{}
	cursor, err := config.Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []primitive.M
	for cursor.Next(context.Background()) {
		var profile bson.M
		err := cursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)
	}
	return Profiles
}

func Getdatabyid(objectid string) []primitive.M {
	id, err := primitive.ObjectIDFromHex(objectid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	cursor, err := config.Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []primitive.M
	for cursor.Next(context.Background()) {
		var profile bson.M
		err := cursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)
	}
	return Profiles
}
