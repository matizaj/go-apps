package main

import (
	"fmt"
	"github.com/matizaj/go-apps/go-breeders/pets"
	"github.com/matizaj/go-apps/go-breeders/utils"
	"net/http"
)

func (app *application) handleHello(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) handlePage(w http.ResponseWriter, r *http.Request) {
	page := r.PathValue("page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) createDogFromFactory(w http.ResponseWriter, r *http.Request) {
	dog := pets.NewPet("dog")
	utils.WriteJson(w, http.StatusOK, dog)
}
func (app *application) createCatFromFactory(w http.ResponseWriter, r *http.Request) {
	dog := pets.NewPet("cat")
	utils.WriteJson(w, http.StatusOK, dog)
}
func (app *application) testPattern(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}
func (app *application) createDogFromAbsFac(w http.ResponseWriter, r *http.Request) {
	dog, err := pets.NewPetFromAbstractFactory("dog")
	fmt.Println("dog", dog)
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, dog)
}
func (app *application) createDogFromAbs(w http.ResponseWriter, r *http.Request) {
	dog, err := pets.NewPetFromAbstractFactory("dog")
	fmt.Println("dog", dog)
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, dog)
}
func (app *application) createCatFromAbsFac(w http.ResponseWriter, r *http.Request) {
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, cat)
}
func (app *application) getAllDogBreeds(w http.ResponseWriter, r *http.Request) {
	breeds, err := app.App.Models.DogBreed.GetAll()
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, breeds)
}
func (app *application) createPetWithBuilder(w http.ResponseWriter, r *http.Request) {
	pet := r.PathValue("pet")
	animal, err := pets.NewPetBuilder().
		SetSpecies(pet).
		SetBreed("big one").
		SetWeight(25).
		SetDescription("some kind of description").
		SetColor("brown").
		SetMinWeight(1).
		SetMaxWeight(9).Build()
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}

	utils.WriteJson(w, http.StatusOK, animal)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	cb, err := app.catService.Remote.GetAllCatBreeds()
	if err != nil {
		utils.ErrorJson(w, err, http.StatusInternalServerError)
	}

	utils.WriteJson(w, http.StatusOK, cb)
}
func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	// get speciers from url

	// get breed from url

	// create breed from abs factory

	// write result to Json

	utils.WriteJson(w, http.StatusOK, cb)
}
