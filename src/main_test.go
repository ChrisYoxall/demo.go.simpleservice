package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGreeting(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Error making requst: %s", err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(GetGreeting)
	handler.ServeHTTP(r, req)

	expected := "Hello"

	if r.Body.String() != expected {
		t.Errorf("Unexpected version. Expected %s but got %s", expected, r.Body.String())
	}

}
