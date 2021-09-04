package main

import (
	"context"
	"log"
	"time"
)

func main() {
	// DB Connection
	client, coll, err := connectToDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Could not connect to DB: %v\n", err)
	}
	defer func() {
		ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// initialize server instance
	s := Server{coll}
	s.ListenAndServe()
}
