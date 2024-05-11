package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", app.home)
	router.HandleFunc("GET /cities", app.allCities)
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
func (app *application) allCities(w http.ResponseWriter, r *http.Request) {
	c, err := app.App.Models.City.GetAll()
	if err != nil {
		log.Println(err)
	}
	b, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}
