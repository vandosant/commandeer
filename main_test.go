package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"encoding/json"
)

func TestCommandReturns200(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	w := httptest.NewRecorder()
	CommandHandler(w, req)

	if w.Code != 200 {
		t.Errorf("response code incorrect: Actual %d, Expected: %s", w.Code, 200)
	}
}

func TestCommandReturnsCollectionJsonType(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	w := httptest.NewRecorder()
	CommandHandler(w, req)

	expected := "application/vnd.application+json; charset=UTF-8"

	result := w.HeaderMap.Get("Content-Type")

	if strings.Contains(result, expected) != true {
		t.Errorf("response format incorrect: Actual %s, Expected: %s", result, expected)
	}
}

func TestCommandReturnsCollectionJson(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	w := httptest.NewRecorder()
	CommandHandler(w, req)

	expected := `{"collection":`

	result := w.Body.String()

	if strings.Contains(result, expected) != true {
		t.Errorf("response format incorrect: Actual %s, Expected: %s", result, expected)
	}
}

func TestCommandReturnsCommandCollection(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	w := httptest.NewRecorder()
	CommandHandler(w, req)

	v := struct {
		Collection  string `json:"collection"`
		Command []struct {
			Name string `json:"name"`
		} `json:"commands"`
	}{}

	if err := json.NewDecoder(w.Body).Decode(&v); err != nil {
		t.Errorf("Should be able to unmarshal response.")
	}

	t.Log("Should be able to unmarshal response.")
}
