package bookstore

import (
	"errors"
	"fmt"
)

const (
	CategoryTech = "Tech"
	CategoryLargePrintRomance = "Large Print Romance"
	CategoryParticlePhysics = "Particle Physics"
)

var validCategory = map[string]bool{
	CategoryTech : true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics: true,
}

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        string
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

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("price %d is invalid", price)
	}
	b.PriceCents = price
	return nil
}

func (b *Book) SetCatrgory(category string) error {
	if validCategory[category] {
		return fmt.Errorf("unknown category %q", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() string {
	return b.category
}
