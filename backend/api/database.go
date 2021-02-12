package api

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	URI = "mongodb+srv://root:kalyani123456@cluster0.7qjuo.mongodb.net/jobBazar?retryWrites=true&w=majority"
)

func DatabaseConnect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		URI,
	))

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}
