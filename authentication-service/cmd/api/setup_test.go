package main

import (
	"github.com/matizaj/go-app/authentication-service/data"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	repo := data.NewPostgresTestRepo(nil)
	testApp.Repo = repo
	os.Exit(m.Run())
}
