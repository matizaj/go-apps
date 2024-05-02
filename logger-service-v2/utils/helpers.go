package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ReadJson(r *http.Request, payload any) error {

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("failed to read request body")
		return err
	}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		log.Println("failed to parse request body")
		return err
	}
	return nil
}
