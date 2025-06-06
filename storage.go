package main

import (
	"sync"
)

var (
	quotes []Quote
	nextID = 1
	mu     sync.Mutex
)

func InitializeStorage() {
	quotes = make([]Quote, 0)
}

func GetAllQuotes() []Quote {
	mu.Lock()
	defer mu.Unlock()
	return quotes
}

func AddQuoteToStorage(quote Quote) Quote {
	mu.Lock()
	defer mu.Unlock()

	quote.ID = nextID
	nextID++
	quotes = append(quotes, quote)
	return quote
}

func DeleteQuoteFromStorage(id int) bool {
	mu.Lock()
	defer mu.Unlock()

	for i, quote := range quotes {
		if quote.ID == id {
			quotes = append(quotes[:i], quotes[i+1:]...)
			return true
		}
	}
	return false
}

func FilterQuotesByAuthor(author string) []Quote {
	mu.Lock()
	defer mu.Unlock()

	var filtered []Quote
	for _, quote := range quotes {
		if quote.Author == author {
			filtered = append(filtered, quote)
		}
	}
	return filtered
}
