package bookstore_test

import (
	"bookstore/bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:           "Shago jeun",
		Author:          "La scepe",
		Copies:          5,
		PriceCents:      30,
		DiscountPercent: 0,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Sparkles of Joy",
		Author: "Maria Jones",
		Copies: 5,
	}
	want := 4
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies

	if want != got {
		t.Errorf("want %d copies after buying 1 copy from a stock of 5, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Sparkles of Joy",
		Author: "Maria Jones",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)

	if err == nil {
		t.Error("want error buying book wehn zero copies left, but got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {Title: "Kosere ni Moscow"},
		2: {Title: "Lord of Lambas"},
	}
	want := []bookstore.Book{
		{Title: "Kosere ni Moscow"},
		{Title: "Lord of Lambas"},
	}
	got := catalog.GetAllBooks()

	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "Book One"},
		2: {ID: 2, Title: "Book Two"},
		3: {ID: 3, Title: "Book Three"},
	}
	want := bookstore.Book{ID: 2, Title: "Book Two"}
	got, err := catalog.GetBook(2)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookIdDoesNotExit(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "Book One"},
	}
	_, err := catalog.GetBook(3)
	if err == nil {
		t.Error("want error for invlid ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "Alaroye",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := b.NetPriceCents()
	if want != got {
		t.Errorf("with price %d, after %d%% discount want net %d got %d", b.PriceCents, b.DiscountPercent, want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "Book one",
		PriceCents: 5,
	}
	want := 100
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "Another one",
		PriceCents: 25,
	}
	err := b.SetPriceCents(-5)

	if err == nil {
		t.Error("want error setting price to -5, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	b := bookstore.Book{
		Title: "Sola wanmbe",
	}
	err := b.SetCatrgory(bookstore.CategoryTech)
	if err != nil {
		t.Fatal(err)
	}
	want := "tech"
	got := b.Category()
	if want != got {
		t.Errorf("want category %q, got %q", want, got)
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	b := bookstore.Book{
		Title: "Ilesanmi Health Edu",
	}
	err := b.SetCatrgory("anything")
	if err == nil {
		t.Error("want error setting category anything, got nil")
	}
}