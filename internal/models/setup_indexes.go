package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func CreateUserIndexes() {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	name, err := DB.Collection("users").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal("Can't create Index " + name + " : " + err.Error())
	}

	log.Print("Created index: " + name)
}

func SetupIndexes() {
	CreateUserIndexes()
}
