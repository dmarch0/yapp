package db

import (
	"context"
	"log"
	"time"
	"yapp/core/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongo() {
	config := configs.GetConfig()

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDBConfig.Uri))

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err = client.Ping(pingCtx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
}
