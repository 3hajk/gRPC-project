package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const DBNAME = "test"

func Connect() (client *mongo.Client) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("[-] Error while connecting to the database: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("[-] Ping error: %v", err)
	}
	log.Println("[+] Succesfully connected to the database.")
	return client
}

func Close(client *mongo.Client) {

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Error while closing connection to the database: %v", err)
	}
	log.Println("[*] Connection to the database has been closed.")
}

func InsertProductToTheDB(ctx context.Context, client *mongo.Client, data ProductItem) error {
	collection := client.Database(DBNAME).Collection("products")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"Name", data.Name}}
	query := bson.M{
		"$setOnInsert": bson.M{
			"lastUpdate": data.LastUpdate,
			"price":      data.Price,
		},
		"$set": bson.M{
			"name":       data.Name,
			"price":      data.Price,
			"lastUpdate": data.LastUpdate,
			"count":      data.Count,
		},
		"$inc": bson.M{
			"count": 1,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, query, opts)
	if err != nil {
		return err
	}
	return nil
}
