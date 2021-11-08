package domain

import (
	"errors"
	"time"
)

var (
	// ErrBeerNotFound is used when a beer could not be found
	ErrBeerNotFound = errors.New("beer not found")

	// ErrDuplicateBeer is used when a beer already exists
	ErrDuplicateBeer = errors.New("beer already exists")
)

// Beer defines the properties of a beer
type Beer struct {
	ID        string
	Name      string
	Brewery   string
	Abv       float32
	ShortDesc string
	Created   time.Time
}

type BeerRepository interface {
	GetBeer(ID string) (Beer, error)
	GetBeers() ([]Beer, error)
	AddBeer(b Beer) error
}
