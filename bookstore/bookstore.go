package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
	PriceCents int
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	book, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return book, nil
}

func (b Book) NetPriceCents() int {
	discount := b.DiscountPercent * b.PriceCents / 100
	return b.PriceCents - discount
}