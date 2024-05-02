package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const mongoUrl = "mongodb://mongo-v2:27017"

func ConnectToMongo() (*mongo.Client, error) {
	clientOpt := options.Client().ApplyURI(mongoUrl)
	clientOpt.SetAuth(options.Credential{
		Username: "go-admin",
		Password: "password",
	})

	connection, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		log.Panicf("cant connect to mongo db %s", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	defer func() {
		if err = connection.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = connection.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("cant ping mongo ->", err)
		return nil, err
	}

	return connection, nil
}
