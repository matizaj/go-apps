package main

import (
	"fmt"
	"github.com/matizaj/go-app/log-service-v2/data"
	"github.com/matizaj/go-app/log-service-v2/database"
	"log"
	"net/http"
)

const (
	webPort = "8090"
)

type Config struct {
	Models *data.Models
}

func main() {
	log.Println("Logger-v2")

	mongoClient, err := database.ConnectToMongo()
	if err != nil {
		log.Panicln("mongo connection failed")
	}

	app := Config{}
	app.Models = data.New(mongoClient)

	app.serve()
}

func (app *Config) serve() {
	log.Println("Starting logging service on port ", webPort)

	router := http.NewServeMux()
	app.RegisterRoutes(router)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Panicf("logger service api cant start %s", err)
	}
}
