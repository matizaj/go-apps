package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matizaj/go-app/e-com/config"
	"github.com/matizaj/go-app/e-com/types"
	"github.com/matizaj/go-app/e-com/utils"
	"log"
	"net/http"
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

func WithJwtAuth(handlerFunc http.HandlerFunc, repo types.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from user request
		tokenString := getTokenFromRequest(r)

		// validate jwt
		token, err := validateToken(tokenString)
		if err != nil {
			log.Println("failed to validate jwt token", err)
			permitionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid jwt token", err)
			permitionDenied(w)
			return
		}

		// fetch user id from db
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userId"].(string)
		uid, _ := strconv.Atoi(str)
		u, err := repo.GetUserById(uid)
		if err != nil {
			log.Println("failed to get user id", err)
			permitionDenied(w)
			return
		}

		// set context uid
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userId", u.Id)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}
}
func GetUserIdFromCtx(ctx context.Context) int {
	uid, ok := ctx.Value("userId").(int)
	if !ok {
		return -1
	}
	return uid
}

func permitionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
}

func getTokenFromRequest(r *http.Request) string {
	token := r.Header.Get("token")
	if token == "" {
		return ""
	}
	return token
}
func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(config.Envs.JwtSecret), nil
	})
}
