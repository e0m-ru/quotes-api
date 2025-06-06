package main

import (
	"testing"
)

func TestInitializeStorage(t *testing.T) {
	InitializeStorage()
	if len(quotes) != 0 {
		t.Fatalf("Ожидалось, что хранилище будет пустым, но там %d цитат", len(quotes))
	}
}
func TestAddQuoteToStorage(t *testing.T) {
	InitializeStorage()
	quote := Quote{Author: "Тестовый Автор", Text: "Тестовая цитата"}
	addedQuote := AddQuoteToStorage(quote)
	if addedQuote.ID == 0 {
		t.Fatalf("Цитата не была добавлена, ID равен 0")
	}
	if len(quotes) != 1 {
		t.Fatalf("Ожидалось 1 цитата в хранилище, но там %d", len(quotes))
	}
	if quotes[0].Author != quote.Author || quotes[0].Text != quote.Text {
		t.Fatalf("Добавленная цитата не совпадает с ожидаемой")
	}
}
func TestGetAllQuotes(t *testing.T) {
	InitializeStorage()
	quote1 := Quote{Author: "Автор 1", Text: "Цитата 1"}
	quote2 := Quote{Author: "Автор 2", Text: "Цитата 2"}
	AddQuoteToStorage(quote1)
	AddQuoteToStorage(quote2)
	allQuotes := GetAllQuotes()
	if len(allQuotes) != 2 {
		t.Fatalf("Ожидалось 2 цитаты, получено %d", len(allQuotes))
	}
}
func TestDeleteQuoteFromStorage(t *testing.T) {
	InitializeStorage()
	quote := Quote{Author: "Автор", Text: "Цитата"}
	addedQuote := AddQuoteToStorage(quote)
	deleted := DeleteQuoteFromStorage(addedQuote.ID)
	if !deleted {
		t.Fatalf("Цитата не была удалена")
	}
	if len(quotes) != 0 {
		t.Fatalf("Ожидалось, что хранилище будет пустым, но там %d цитат", len(quotes))
	}
	deletedAgain := DeleteQuoteFromStorage(addedQuote.ID)
	if deletedAgain {
		t.Fatalf("Удаление уже удаленной цитаты должно вернуть false")
	}
}
func TestFilterQuotesByAuthor(t *testing.T) {
	InitializeStorage()
	quote1 := Quote{Author: "Автор 1", Text: "Цитата 1"}
	quote2 := Quote{Author: "Автор 2", Text: "Цитата 2"}
	quote3 := Quote{Author: "Автор 1", Text: "Цитата 3"}
	AddQuoteToStorage(quote1)
	AddQuoteToStorage(quote2)
	AddQuoteToStorage(quote3)
	filteredQuotes := FilterQuotesByAuthor("Автор 1")
	if len(filteredQuotes) != 2 {
		t.Fatalf("Ожидалось 2 цитаты от 'Автор 1', получено %d", len(filteredQuotes))
	}
}
