package main

import (
	models "github.com/matizaj/go-app/log-service-v2/data"
	"github.com/matizaj/go-app/log-service-v2/utils"
	"log"
	"net/http"
)

func (app *Config) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", app.hello)
	router.HandleFunc("POST /log", app.addLogEntry)
}

func (app *Config) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
func (app *Config) addLogEntry(w http.ResponseWriter, r *http.Request) {
	var payload models.LogEntry
	err := utils.ReadJson(r, &payload)
	if err != nil {
		log.Println("failed to parse request", err)
		return
	}
	err = app.Models.Log.AddLogEntry(payload)
	if err != nil {
		log.Println("failed to add log entry", err)
		return
	}
	w.Write([]byte("ok"))
}
