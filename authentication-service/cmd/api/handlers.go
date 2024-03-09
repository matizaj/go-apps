package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Auth(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJson(w, r, &requestPayload)
	log.Printf(" received payload %s: ", requestPayload)
	if err != nil {
		log.Printf("cant read payload %s: ", err)
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	users, err := app.Models.User.GetAll()
	log.Printf("user %s: ", user)
	if err != nil {
		log.Printf("cant find user %s: ", err)
		errors.New("invalid credentials")
		return
	}
	log.Printf("user %s: ", user)
	log.Printf("users %s: ", users)
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		errors.New("invalid credentials")
		return
	}
	log.Printf("valid credsd %s: ", valid)

	// log authentication
	log.Printf("start log auth request...")
	err = app.logRequest("auth-log", "user email addr")
	if err != nil {
		log.Printf("cant log auth request... ", err)
		errors.New("cant log auth request")
		return
	}
	log.Printf("finished log auth request...")
	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}
	log.Printf("payload to send %s: ", payload)
	app.writeJson(w, http.StatusOK, payload)
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
	default:
		app.errorJson(w, errors.New("unknown action"), http.StatusBadRequest)
	}
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

func (app *Config) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, err := json.MarshalIndent(entry, "", "\t")
	if err != nil {
		log.Println("cant pars data ", err)
		return err
	}

	logServiceUrl := "http://logger-service/log"
	request, err := http.NewRequest("POST", logServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("failed to send data top logger service.. ", err)
		return err
	}
	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Println("failed to get response from logger service.. ", err)
		return err
	}
	defer resp.Body.Close()
	log.Printf("request logged")
	return nil
}
