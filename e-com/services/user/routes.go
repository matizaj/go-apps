package user

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
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

}
