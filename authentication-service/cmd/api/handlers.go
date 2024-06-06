package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	user, err := app.Repo.GetByEmail(requestPayload.Email)
	if err != nil {
		log.Printf("cant find user %s: ", err)
		errors.New("invalid credentials")
		return
	}
	valid, err := app.Repo.PasswordMatches(requestPayload.Password, *user)
	if err != nil || !valid {
		errors.New("invalid credentials")
		return
	}

	// log authentication
	log.Printf("start log auth request...")
	err = app.logRequest("auth-log", "user email addr")
	if err != nil {
		errors.New("cant log auth request")
		return
	}
	log.Printf("finished log auth request...")
	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}
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

	log.Println("send  auth request")
	response, err := app.Client.Do(req)
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
func (app *Config) HandleRegister(w http.ResponseWriter, req *http.Request) {
	type user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var u user

	defer req.Body.Close()

	body, _ := io.ReadAll(req.Body)
	_ = json.Unmarshal(body, &u)

	buser, _ := json.Marshal(u)
	w.Write(buser)
}
