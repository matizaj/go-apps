package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func Test_Auth(t *testing.T) {
	jsonToReturn := `
		{
			"error":false,
			"message":"fake response"
		}
	`
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	testApp.Client = client

	postBody := map[string]interface{}{
		"email":    "emial@test.com",
		"password": "secret",
	}

	body, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/auth", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.Auth)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 but got %d ", rr.Code)
	}
}
