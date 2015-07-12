package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexReturnsText(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	IndexHandler(recorder, req)

	expected := "Welcome!"

	result := recorder.Body.String()

	if strings.Contains(result, expected) != true {
		t.Errorf("json format incorrect: Actual %s, Expected: %s", result, expected)
	}
}

func TestCmdsReturns200(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/cmds", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	w := httptest.NewRecorder()
	CommandHandler(w, req)

	if w.Code != 200 {
		t.Errorf("response code incorrect: Actual %d, Expected: %s", w.Code, 200)
	}
}
