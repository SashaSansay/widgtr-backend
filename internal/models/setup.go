package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

var DB *mongo.Database
var MongoClient *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoUrl := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}

	MongoClient = client
	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "widgtr"
	}
	DB = client.Database(dbName)

	SetupIndexes()
}

func DisconnectDB() {
	if err := MongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
