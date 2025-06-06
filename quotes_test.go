package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddQuote(t *testing.T) {
	InitializeStorage()

	payload := `{"author": "Test Author", "quote": "Test Quote"}`
	req, err := http.NewRequest("POST", "/quotes", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddQuote)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	expected := `"author":"Test Author"`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetQuotes(t *testing.T) {
	InitializeStorage()
	AddQuoteToStorage(Quote{Author: "Test Author", Text: "Test Quote"})

	req, err := http.NewRequest("GET", "/quotes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetQuotes)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `"author":"Test Author"`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}