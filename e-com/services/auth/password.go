package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(plain string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
