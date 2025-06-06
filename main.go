package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Инициализация хранилища
	InitializeStorage()

	// Маршруты API
	router.HandleFunc("/quotes", GetQuotes).Methods("GET")
	router.HandleFunc("/quotes", AddQuote).Methods("POST")
	router.HandleFunc("/quotes/random", GetRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", DeleteQuote).Methods("DELETE")

	// Запуск сервера
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
