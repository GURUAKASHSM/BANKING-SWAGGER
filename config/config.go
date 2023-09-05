package config

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectstring = "mongodb+srv://guru:guru@banking.sy1piq8.mongodb.net/?retryWrites=true&w=majority"
const dbname = "Banking"
const dbcollection = "Coustomer_transaction"

var Collection *mongo.Collection

func init() {
	clientoption := options.Client().ApplyURI(connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	Collection = client.Database(dbname).Collection(dbcollection)

	fmt.Println("Collection Connected")
}
