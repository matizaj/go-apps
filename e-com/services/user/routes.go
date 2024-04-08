package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/matizaj/go-app/e-com/config"
	"github.com/matizaj/go-app/e-com/services/auth"
	"github.com/matizaj/go-app/e-com/types"
	"github.com/matizaj/go-app/e-com/utils"
	"log"
	"net/http"
)

type Handler struct {
	repository types.UserRepository
}

func NewHandler(repo types.UserRepository) *Handler {
	return &Handler{repository: repo}
}

func (h *Handler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("GET /", h.handleHello)
	router.HandleFunc("POST /login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)
}

func (h *Handler) handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("logging user")
	var payload types.LoginUserPayload
	err := utils.ParseJson(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.Validate.Struct(payload)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	user, err := h.repository.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, 500, fmt.Errorf("user %s not exists", payload.Email))
		return
	}
	credsValid, err := auth.ComparePassword(user.Password, payload.Password)
	if err != nil || !credsValid {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	token, err := auth.CreateJwt([]byte(config.Envs.JwtSecret), user.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	log.Println(" registering...")
	// get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	// check if the user exist
	_, err := h.repository.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, 500, fmt.Errorf("user %s already exists", payload.Email))
		return
	}
	// create new user or drop req
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("hash password failed"))
		return
	}
	log.Println("b4 c u ")
	err = h.repository.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, 500, fmt.Errorf("cant create user"))
		return
	}
	utils.WriteJson(w, http.StatusCreated, nil)
}
