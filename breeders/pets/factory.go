package pets

import "github.com/matizaj/go-apps/go-breeders/models"

func NewPet(species string) *models.Pet {
	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no description yet",
		LifeSpan:    0,
	}
	return &pet
}
