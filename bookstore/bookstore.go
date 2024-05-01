package bookstore

import (
	"errors"
	"fmt"
	"sort"
)

// Book represents information about a book.
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}

func NetPriceCents(b Book) int {
	return b.PriceCents - (b.DiscountPercent * b.PriceCents / 100)
}
