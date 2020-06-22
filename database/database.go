package database

import (
	"context"
	pb "github.com/3hajk/grpc-project/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const DBNAME = "test"

func Connect() (client *mongo.Client) {

	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
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

func SetIndex(client *mongo.Client) {
	index := mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	collection := client.Database(DBNAME).Collection("products")
	_, err := collection.Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		log.Fatal(err)
	}
}

func Close(client *mongo.Client) {

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Error while closing connection to the database: %v", err)
	}
	log.Println("[*] Connection to the database has been closed.")
}

func InsertProductToDB(ctx context.Context, client *mongo.Client, data ProductItem) error {
	collection := client.Database(DBNAME).Collection("products")
	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"name": data.Name,
	}

	query := bson.M{
		"$setOnInsert": bson.M{
			"name": data.Name,
		},
		"$set": bson.M{
			"price":      data.Price,
			"lastUpdate": data.LastUpdate,
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

func GetProductFromDB(ctx context.Context, client *mongo.Client, skip int64, limit int64, sortName string, sortDirect int32) ([]*pb.Product, error) {
	collection := client.Database(DBNAME).Collection("products")
	opts := &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort:  bson.M{sortName: sortDirect},
	}
	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Println("Couldn't get sugar levels", "err", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	results := make([]*pb.Product, 0, 0)
	for cursor.Next(context.Background()) {
		var result = new(ProductItem)
		err := cursor.Decode(result)
		if err != nil {
			break
		}
		results = append(results, &pb.Product{
			Name:        result.Name,
			Price:       float32(result.Price),
			LastUpdate:  result.LastUpdate.T,
			ChangePrice: int32(result.Count),
		})
	}

	return results, nil
}
