package main

import (
	"context"
	"log"
	"os"
	"time"
)

func main() {
	// DB Connection
	client, coll, err := connectToDB(os.Getenv("MONGO_URL"))
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
	s := Server{coll, nil}
	s.ListenAndServe()
}
