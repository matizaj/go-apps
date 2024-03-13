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
func Contact(w http.ResponseWriter, r *http.Request) {
	reqId := r.PathValue("id")

	fmt.Println("req id: ", reqId)
	if reqId != "" {
		fmt.Println("req id: ", reqId)
	}
	fmt.Fprintf(w, "contact-page")
}

func middleware(next http.Handler) http.Handler {
	fmt.Println("hit middleware 1")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit middleware 2")
		next.ServeHTTP(w, r)
	})
}
func main() {

	fmt.Println("...main...")
	http.Handle("/", middleware)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact/{id}", Contact)
	http.HandleFunc("/contact/", Contact)
	err := http.ListenAndServe(":4041", nil)
	fmt.Println("Server is listening on port :4040")
	if err != nil {
		log.Panic("server cant start")
	}

}
