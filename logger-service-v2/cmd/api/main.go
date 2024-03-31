package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

const (
	mongoUrl = "mongodb://localhost:27017"
	webPort  = "8090"
)

var client *mongo.Client

type Config struct{}

func main() {
	log.Println("Logger-v2")
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panicln("mongo connection failed")
	}
	client = mongoClient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	app := Config{}
	app.serve()
}

func connectToMongo() (*mongo.Client, error) {
	clientOpt := options.Client().ApplyURI(mongoUrl)
	clientOpt.SetAuth(options.Credential{
		Username: "go-admin",
		Password: "password",
	})

	connection, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		log.Panicf("cant connect to mopngo db %s", err)
		return nil, err
	}
	return connection, nil
}

func (app *Config) serve() {
	log.Println("Starting logging service on port ", webPort)
	http.HandleFunc("/", app.Hello)
	err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), nil)
	if err != nil {
		log.Panicf("logger service cant start %s", err)
	}

}
