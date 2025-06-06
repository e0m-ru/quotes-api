package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")

	var quotes []Quote
	if author != "" {
		quotes = FilterQuotesByAuthor(author)
	} else {
		quotes = GetAllQuotes()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func AddQuote(w http.ResponseWriter, r *http.Request) {
	var newQuote Quote
	err := json.NewDecoder(r.Body).Decode(&newQuote)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addedQuote := AddQuoteToStorage(newQuote)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(addedQuote)
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quotes := GetAllQuotes()
	if len(quotes) == 0 {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}

	randomQuote := quotes[rand.Intn(len(quotes))]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(randomQuote)
}

func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid quote ID", http.StatusBadRequest)
		return
	}

	if !DeleteQuoteFromStorage(id) {
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
