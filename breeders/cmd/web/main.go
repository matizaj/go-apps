package main

import (
	"flag"
	"fmt"
	"github.com/matizaj/go-apps/go-breeders/configuration"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
	catService  *RemoteService
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}
	flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
	flag.StringVar(&app.config.dsn, "dns", "breeders:mypassword@tcp(localhost:3336)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	//get database
	db, err := initMySqlDb(app.config.dsn)
	if err != nil {
		log.Panicln(err)
	}

	//jsonBackend := &jsonBackend{}
	//jsonAdapter := &RemoteService{Remote: jsonBackend}

	xmlBackend := &xmlBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	app.App = configuration.New(db)
	app.catService = xmlAdapter

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

	err = server.ListenAndServe()
	if err != nil {
		log.Panicln("server failed to start", err)
	}

	fmt.Println("Starting web application on port ", port)
}
