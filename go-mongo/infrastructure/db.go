package infrastructure

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	clientOptions.SetAuth(options.Credential{
		Username: "matt",
		Password: "password",
	})
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Connection error")
		return nil, err
	}
	return c, nil
}
