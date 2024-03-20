package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"test/go-mongo/infrastructure"
	models2 "test/go-mongo/models"
	"time"
)

var client *mongo.Client

func insertItem(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("docker").Collection("users")
	insert, err := collection.InsertOne(context.TODO(), models2.User{Name: "Hania", Age: 4})
	if err != nil {
		http.Error(w, "failed to insert data", http.StatusInternalServerError)
		return
	}
	log.Println("insertion id: ", insert.InsertedID)
	w.Write([]byte("ok"))
}

func findOne(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("docker").Collection("users")
	filter := bson.D{{"name", "Hania"}}
	var result models2.User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("failed to find one")
	}
	log.Println("result", result)
	w.Write([]byte(result.Name))
}

func main() {
	mongoClient, err := infrastructure.ConnectToMongo()

	if err != nil {
		log.Panicln(err)
	}
	client = mongoClient
	// create context to disconnect mongo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	http.HandleFunc("/insert", insertItem)
	http.HandleFunc("/find", findOne)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Panicln("App crashed!")
	}
}
