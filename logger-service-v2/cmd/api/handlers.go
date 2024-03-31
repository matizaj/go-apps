package main

import "net/http"

func (app *Config) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
