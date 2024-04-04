package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/matizaj/go-app/e-com/services/auth"
	"github.com/matizaj/go-app/e-com/types"
	"github.com/matizaj/go-app/e-com/utils"
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

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}
	// check if the user exist
	_, err := h.repository.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, 500, fmt.Errorf("something went wrong"))
		return
	}
	// create new user or drop req
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("hash password failed"))
		return
	}
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
