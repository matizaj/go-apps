package main

import (
	"github.com/matizaj/go-apps/go-breeders/configuration"
	"github.com/matizaj/go-apps/go-breeders/models"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &TestBackend{}
	testAdapter := &RemoteService{Remote: testBackend}

	testApp = application{
		App:        configuration.New(nil),
		catService: testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{Id: 22, Breed: "Tomcat", Details: "some details"},
	}

	return breeds, nil
}
