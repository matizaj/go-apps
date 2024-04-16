package main

import (
	"github.com/matizaj/go-apps/go-breeders/configuration"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {

	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
