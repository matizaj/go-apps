package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type Application struct {
	templateMap map[string]*template.Template
	config      appConfig
}

type appConfig struct {
	useCache bool
}

func main() {
	app := Application{
		templateMap: make(map[string]*template.Template),
	}
	flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
	flag.Parse()

	router := http.NewServeMux()
	app.RegisterRoute(router)

	server := &http.Server{
		Addr:              port,
		Handler:           router,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panicln("server failed to start", err)
	}

	fmt.Println("Starting web application on port ", port)
}
