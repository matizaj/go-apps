package main

import (
	"github.com/matizaj/go-app/log-service/data"
	"log"
	"net/http"
)

type JsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// TODO
	// handle insert log to db
	var reqPayload JsonPayload
	err := app.readJson(w, r, &reqPayload)
	if err != nil {
		log.Println("cant parse request payload: ", err)
	}

	event := data.LogEntry{
		Name: reqPayload.Name,
		Data: reqPayload.Data,
	}

	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		log.Println("cant insert log event to database: ", err)
	}

	resp := jsonResponse{
		Error:   false,
		Data:    event,
		Message: "logged",
	}

	app.writeJson(w, http.StatusOK, resp)
}
