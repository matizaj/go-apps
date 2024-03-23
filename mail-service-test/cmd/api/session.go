package main

import (
	"errors"
	"fmt"
	"github.com/matizaj/go-app/mail-service/data/models"
	"net/http"
)

func getUser(r *http.Request) models.User {
	var u models.User
	cookie, err := r.Cookie("sid")
	if errors.Is(http.ErrNoCookie, err) {
		return u
	}

	// user arleady exist
	if userName, ok := dbSession[cookie.Value]; ok {
		u = dbUsers[userName]
	}

	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("sid")
	if errors.Is(err, http.ErrNoCookie) {
		return false
	}
	userName, _ := dbSession[cookie.Value]
	fmt.Println("username ", userName)
	_, ok := dbUsers[userName]
	return ok
}
