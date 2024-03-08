package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "hit the broker",
	}
	_ = app.writeJson(w, http.StatusOK, payload, nil)
}
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var reqPayload RequestPayload
	err := app.readJson(w, r, &reqPayload)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}
	switch reqPayload.Action {
	case "auth":
		app.authenticate(w, reqPayload.Auth)
	case "log":
		app.logItem(w, reqPayload.Log)
	default:
		app.errorJson(w, errors.New("unknown action"), http.StatusBadRequest)
	}
}

func (app *Config) logItem(w http.ResponseWriter, entry LogPayload) {
	jsonData, err := json.MarshalIndent(entry, "", "\t")
	if err != nil {
		log.Println("failed to read payload")
		return
	}
	logServiceUrl := "http://logger-service/log"
	req, err := http.NewRequest("/log", logServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJson(w, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		app.errorJson(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"
	app.writeJson(w, http.StatusOK, payload)
}
func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {

	// create json to send to the auth service
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	req, err := http.NewRequest("POST", "http://authentication-service/auth", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJson(w, err)
		return
	}
	client := &http.Client{}
	log.Println("send  auth request")
	response, err := client.Do(req)
	log.Printf("response: %s", response)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJson(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		log.Println(response)
		log.Println("something went wrong")
		app.errorJson(w, errors.New("error calling auth service"))
		return
	}

	var jsonFromService jsonResponse
	dec := json.NewDecoder(response.Body)
	log.Println(dec)
	err = dec.Decode(&jsonFromService)
	if err != nil {
		log.Printf("failed to decode json: %s", err)
		app.errorJson(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJson(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromService.Data
	log.Println("parse response")
	app.writeJson(w, 200, payload)
}
