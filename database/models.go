package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductItem struct {
	Name       string              `bson:"name"`
	Price      float64             `bson:"price"`
	LastUpdate primitive.Timestamp `bson:"lastUpdate,omitempty"`
	Count      int                 `bson:"count"`
}
