package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDbClient() *mongo.Client {

	clientOptions := options.Client().ApplyURI("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// logger.Fatal("Error while connecting to DB : " + err.Error())
		panic(err)
	} else {
		// logger.Info("Database Connected Successfully......")

	}
	return client
}
