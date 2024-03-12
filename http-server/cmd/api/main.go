package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type server struct {
}
type Req struct {
	Type string `json:"type"`
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("some shit ", err)
		return
	}
	fmt.Println("type: ", req.Type, r.Method, r.ContentLength)
	w.Header().Set("mateusz", "gomasster")
	w.WriteHeader(203)
	w.Write([]byte("hello"))
}
func main() {
	var s server
	err := http.ListenAndServe(":4040", s)
	if err != nil {
		log.Panicln("server failed to start")
	}
}
