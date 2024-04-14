package pets

import (
	"errors"
	"fmt"
	"github.com/matizaj/go-apps/go-breeders/models"
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
}

type DogAbstractFactory struct {
}

func (df *DogAbstractFactory) newPet() Animal {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type CatAbstractFactory struct {
}

func (cf *CatAbstractFactory) newPet() Animal {
	return &CatFromFactory{
		Pet: &models.Cat{},
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
