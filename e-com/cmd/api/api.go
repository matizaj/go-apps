package api

import (
	"database/sql"
	"github.com/matizaj/go-app/e-com/services/user"
	"log"
	"net/http"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

	userRepo := user.NewStore(s.db)
	userHandler := user.NewHandler(userRepo)
	userHandler.RegisterRoute(router)

	log.Println("Server listening on port 8099")
	err := http.ListenAndServe(s.addr, router)
	if err != nil {
		log.Println("hewre", err)
	}
	return err
}
