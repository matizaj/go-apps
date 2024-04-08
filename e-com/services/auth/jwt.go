package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/matizaj/go-app/e-com/config"
	"strconv"
	"time"
)

func CreateJwt(secret []byte, userId int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JwtExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     strconv.Itoa(userId),
		"expirtedAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
