package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(plain string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func ComparePassword(hash, plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		log.Println("invalid password")
		return false, err
	}
	return true, nil
}
