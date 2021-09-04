package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// connectToDB initializes connection to DB and returns resulting session
func connectToDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	// ping server to make sure connection is alive
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	return client, err
}

// getProductCollection is helper for getting product collection.
// mongo.Client and mongo.Collection are thread safe
func getProductCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("products").Collection("products")
}
