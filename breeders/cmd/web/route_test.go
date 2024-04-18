package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllBreeds(t *testing.T) {
	// create request
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)
	// create response recorder
	rr := httptest.NewRecorder()
	// create a handler
	handler := http.HandlerFunc(testApp.getAllDogBreeds)
	//serve the handler
	handler.ServeHTTP(rr, req)
	// check response
	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code, got %d", rr.Code)
	}
}
func TestApplication_GetAllCatBreeds(t *testing.T) {
	// create request
	req, _ := http.NewRequest("GET", "/api/cat-breeds", nil)
	// create response recorder
	rr := httptest.NewRecorder()
	// create a handler
	handler := http.HandlerFunc(testApp.GetAllCatBreeds)
	//serve the handler
	handler.ServeHTTP(rr, req)
	// check response
	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code, got %d", rr.Code)
	}
}
