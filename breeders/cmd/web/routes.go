package main

import (
	"fmt"
	"net/http"
	"os"
)

//type Handler struct {
//}
//
//func NewHandler() *Handler {
//	return &Handler{}
//}

func (app *Application) RegisterRoute(router *http.ServeMux) {
	directoryPath := "./static"
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", directoryPath)
		return
	}
	fileServer := http.FileServer(http.Dir(directoryPath))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	router.HandleFunc("GET /home", app.handleHello)
	router.HandleFunc("GET /{page}", app.handlePage)

}
func (app *Application) handleHello(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *Application) handlePage(w http.ResponseWriter, r *http.Request) {
	page := r.PathValue("page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}
