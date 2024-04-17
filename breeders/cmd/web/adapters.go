package main

import (
	"encoding/json"
	"github.com/matizaj/go-apps/go-breeders/models"
	"io"
	"log"
	"net/http"
)

type CatBreedsInterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
}

type RemoteService struct {
	Remote CatBreedsInterface
}

func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}

type jsonBackend struct {
}

func (jb *jsonBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	response, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	var catBreeds []*models.CatBreed
	err = json.Unmarshal(body, &catBreeds)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return catBreeds, nil
}
