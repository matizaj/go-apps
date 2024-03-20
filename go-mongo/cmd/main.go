package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Person struct {
	Name string
	Age  string
}

func main() {
	log.Println("App is starting")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions.SetAuth(options.Credential{
		Username: "matt",
		Password: "password",
	})
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Connection error")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Ping failed")
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("test-a")
	fmt.Println("collection created: ", collection.Name())
	insert, err := collection.InsertOne(context.TODO(), Person{
		Name: "matt",
		Age:  "37",
	})
	if err != nil {
		log.Println("failed to insert", err)
	}
	fmt.Println("insert", insert.InsertedID)
	fmt.Println(collection.Name())
	filter := bson.D{{"name", "matt"}}
	var result Person
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("failed to find one")
	}
	log.Println("result", result)

}
