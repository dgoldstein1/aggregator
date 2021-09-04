package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// connectToDB initializes connection to DB and returns resulting session
func connectToDB(uri string) (*mongo.Client, *mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, err
	}
	// ping server to make sure connection is alive
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}
	// do a sample query
	coll := getProductCollection(client)
	opts := options.EstimatedDocumentCount().SetMaxTime(2 * time.Second)
	count, err := coll.EstimatedDocumentCount(context.TODO(), opts)
	if err != nil {
		return nil, nil, err
	}
	log.Printf("Successfully connected to DB. Document count: %v", count)
	return client, coll, err
}

// getProductCollection is helper for getting product collection.
// mongo.Client and mongo.Collection are thread safe
func getProductCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("products").Collection("products")
}
