package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductItem struct {
	ID         string              `bson:"_id, omitempty"`
	Name       string              `bson:"name"`
	Price      float64             `bson:"price"`
	LastUpdate primitive.Timestamp `bson:"lastUpdate,omitempty"`
	Count      uint32              `bson:"count"`
}
