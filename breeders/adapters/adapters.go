package adapters

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
	GetCatBreedByName(b string) (*models.CatBreed, error)
}

type RemoteService struct {
	Remote CatBreedsInterface
}

func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}

type JsonBackend struct {
}
type XmlBackend struct {
}

func (jb *JsonBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
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

func (jb *JsonBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	log.Println("GetCatBreedByName")
	res, err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/json")
	log.Println("response", res)
	if err != nil {
		log.Println("request err", err)
		return nil, err
	}

	var catBreed models.CatBreed
	defer res.Body.Close()

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(bs, &catBreed)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &catBreed, nil
}

func (xb *XmlBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
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
func (xb *XmlBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	res, err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/xml")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	var cb models.CatBreed
	err = xml.Unmarshal(body, &cb)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &cb, nil
}

type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{Id: 22, Breed: "Tomcat", Details: "some details"},
	}

	return breeds, nil
}
func (tb *TestBackend) GetCatBreedByName(b string) (*models.CatBreed, error) {
	return nil, nil
}
