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

func (app *application) RegisterRoute(router *http.ServeMux) {
	directoryPath := "./static"
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", directoryPath)
		return
	}
	fileServer := http.FileServer(http.Dir(directoryPath))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	router.HandleFunc("GET /test-pattern", app.testPattern)
	router.HandleFunc("GET /api/dog-from-factory", app.createDogFromFactory)
	router.HandleFunc("GET /api/cat-from-factory", app.createCatFromFactory)
	router.HandleFunc("GET /api/dog-from-afactory", app.createDogFromAbsFac)
	router.HandleFunc("GET /api/dog-from-abs", app.createDogFromAbs)
	router.HandleFunc("GET /api/cat-from-afactory", app.createCatFromAbsFac)
	router.HandleFunc("GET /api/{pet}", app.createPetWithBuilder)
	router.HandleFunc("GET /api/dog-breeds", app.getAllDogBreeds)
	router.HandleFunc("GET /api/cat-breeds", app.GetAllCatBreeds)
	router.HandleFunc("GET /api/animal-from-abstract-factory/{species}/{breed}", app.GetAnimalFromAbstractFactory)
	router.HandleFunc("GET /home", app.handleHello)
	router.HandleFunc("GET /{page}", app.handlePage)

}
