package main

import (
	"design-patterns/configuration"
	"design-patterns/database"
	"log"
	"net/http"
)

const port = ":7000"

type application struct {
	dsn string
	App *configuration.Application
}

func main() {
	router := http.NewServeMux()
	app := application{
		dsn: "admin:mypassword@tcp(localhost:3366)/population",
	}
	db, err := database.InitDb(app.dsn)
	if err != nil {
		log.Panicln("[From main] failed to init db", err)
	}
	app.App = configuration.New(db)
	app.registerRoutes(router)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panicln("server failed to start")
	}
	log.Println("Starting web application on port ", port)
}
