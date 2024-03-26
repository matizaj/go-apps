package main

import (
	"log"
	"net/http"
)

type mailMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (app *Config) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
func (app *Config) sendMail(w http.ResponseWriter, r *http.Request) {

	var requestPayload mailMessage
	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	log.Println("sendenamil json", requestPayload)
	log.Println("b4 msg")
	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}
	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "send to " + requestPayload.To,
	}

	app.writeJson(w, http.StatusOK, payload)
}
