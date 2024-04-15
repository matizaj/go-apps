package main

import (
	"github.com/matizaj/go-apps/go-breeders/models"
	"os"
	"testing"
)

var testApp Application

func TestMain(m *testing.M) {

	testApp = Application{
		Models: models.NewTest(),
	}

	os.Exit(m.Run())
}
