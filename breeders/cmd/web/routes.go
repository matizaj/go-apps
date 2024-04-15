package main

import (
	"fmt"
	"github.com/matizaj/go-apps/go-breeders/pets"
	"github.com/matizaj/go-apps/go-breeders/utils"
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
	router.HandleFunc("GET /test-pattern", app.testPattern)
	router.HandleFunc("GET /api/dog-from-factory", app.createDogFromFactory)
	router.HandleFunc("GET /api/cat-from-factory", app.createCatFromFactory)
	router.HandleFunc("GET /api/dog-from-afactory", app.createDogFromAbsFac)
	router.HandleFunc("GET /api/dog-from-abs", app.createDogFromAbs)
	router.HandleFunc("GET /api/cat-from-afactory", app.createCatFromAbsFac)
	router.HandleFunc("GET /api/dog-breeds", app.getAllDogBreeds)
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

func (app *Application) createDogFromFactory(w http.ResponseWriter, r *http.Request) {
	dog := pets.NewPet("dog")
	utils.WriteJson(w, http.StatusOK, dog)
}
func (app *Application) createCatFromFactory(w http.ResponseWriter, r *http.Request) {
	dog := pets.NewPet("cat")
	utils.WriteJson(w, http.StatusOK, dog)
}
func (app *Application) testPattern(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}
func (app *Application) createDogFromAbsFac(w http.ResponseWriter, r *http.Request) {
	dog, err := pets.NewPetFromAbstractFactory("dog")
	fmt.Println("dog", dog)
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, dog)
	app.render(w, "test.page.gohtml", nil)
}
func (app *Application) createDogFromAbs(w http.ResponseWriter, r *http.Request) {
	dog, err := pets.NewPetFromAbstractFactory("dog")
	fmt.Println("dog", dog)
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, dog)
}
func (app *Application) createCatFromAbsFac(w http.ResponseWriter, r *http.Request) {
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, cat)
	app.render(w, "test.page.gohtml", nil)
}
func (app *Application) getAllDogBreeds(w http.ResponseWriter, r *http.Request) {
	breeds, err := app.Models.DogBreed.GetAll()
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, breeds)

}
