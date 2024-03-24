package main

import "net/http"

func (app *Config) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
