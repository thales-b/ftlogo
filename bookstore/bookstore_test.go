package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("had %d copies, want %d left, got %d left", b.Copies, want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {
			Title: "For the Love of Go",
			ID:    1,
		},
		2: {
			Title: "The Power of Go: Tools",
			ID:    2,
		},
	}
	want := []bookstore.Book{
		{
			Title: "For the Love of Go",
			ID:    1,
		},
		{
			Title: "The Power of Go: Tools",
			ID:    2,
		},
	}
	got := catalog.GetAllBooks()
	// Seems to always come out sorted
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {
			Title: "For the Love of Go",
		},
		2: {
			Title: "The Power of Go: Tools",
		},
	}
	want := bookstore.Book{
		Title: "For the Love of Go",
	}
	got, err := catalog.GetBook(1)
	if err != nil {
		t.Fatal()
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	want := 1500
	b := bookstore.Book{
		PriceCents:      2000,
		DiscountPercent: 25,
	}
	got := b.NetPriceCents()
	if got != want {
		t.Errorf("want %d, got %d, for price %d and discount %d", want, got, b.PriceCents, b.DiscountPercent)
	}
}
