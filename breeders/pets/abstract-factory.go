package pets

import (
	"errors"
	"fmt"
	"github.com/matizaj/go-apps/go-breeders/configuration"
	"github.com/matizaj/go-apps/go-breeders/models"
	"log"
)

type Animal interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("this animal is a %s", dff.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("this animal is a %s", cff.Pet.Breed.Breed)
}

type PetFactory interface {
	newPet() Animal
	newPetWithBreed(breed string) Animal
}

type DogAbstractFactory struct {
}

func (df *DogAbstractFactory) newPet() Animal {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}
func (df *DogAbstractFactory) newPetWithBreed(breed string) Animal {
	app := configuration.GetInstance()
	b, _ := app.Models.DogBreed.GetBreedByName(breed)
	log.Println("*b", *b)
	log.Println("b", b)
	return &DogFromFactory{
		Pet: &models.Dog{
			Breed: *b,
		},
	}
}

type CatAbstractFactory struct {
}

func (cf *CatAbstractFactory) newPet() Animal {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}
func (cf *CatAbstractFactory) newPetWithBreed(b string) Animal {
	app := configuration.GetInstance()
	breed, _ := app.CatService.Remote.GetCatBreedByName(b)
	log.Println("brred", breed)
	return &CatFromFactory{
		Pet: &models.Cat{Breed: *breed},
	}
}
func NewPetFromAbstractFactory(species string) (Animal, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, errors.New("invalid species")
	}
}

func NewPetFromBreedFromAbstractFactory(species, breed string) (Animal, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPetWithBreed(breed)
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPetWithBreed(breed)
		log.Println("cat", cat)
		return cat, nil

	default:
		return nil, errors.New("invalid species")
	}
}
