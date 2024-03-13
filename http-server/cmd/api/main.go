package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBody struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func About(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("cant parse request payload")
	}
	data, err := json.Marshal(requestBody)
	fmt.Println("body: ", requestBody)
	w.Header().Set("my-custom-header", "matesz-zajac")
	w.Write(data)
}
func main() {

	fmt.Println("...main...")

	http.HandleFunc("/about", About)
	err := http.ListenAndServe(":4041", nil)
	fmt.Println("Server is listening on port :4040")
	if err != nil {
		log.Panic("server cant start")
	}

}
