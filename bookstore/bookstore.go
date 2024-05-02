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
	category        Category
}

// Catalog represents a collection of books grouped by ID.
type Catalog map[int]Book

type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (catalog Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func (catalog Catalog) GetBook(ID int) (Book, error) {
	b, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}

func (b Book) NetPriceCents() int {
	return b.PriceCents - (b.DiscountPercent * b.PriceCents / 100)
}

func (b *Book) SetPriceCents(new int) error {
	if new < 0 {
		return fmt.Errorf("negative price %d", new)
	}
	b.PriceCents = new
	return nil
}

func (b Book) Category() Category {
	return b.category
}

func (b *Book) SetCategory(category Category) error {
	if !validCategory[category] {
		return fmt.Errorf("unknown category %v", category)
	}
	b.category = category
	return nil
}
