package main

import (
	"encoding/json"
	"encoding/xml"
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
type xmlBackend struct {
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

func (xb *xmlBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	response, err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	type catBreeds struct {
		XMLName struct{}           `xml:"cat-breeds"`
		Breeds  []*models.CatBreed `xml:"cat-breed"`
	}
	var cb catBreeds
	err = xml.Unmarshal(body, &cb)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return cb.Breeds, nil
}
