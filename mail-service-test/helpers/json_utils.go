package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/matizaj/go-app/mail-service/data/models"
	"io"
	"net/http"
)

func ReadJson(r *http.Request) (*models.User, error) {

	var user models.User
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil, err
	}
	fmt.Println("USER: ", user)
	return &user, nil
}
