package service

import (
	"github.com/katzien/go-structure-examples/new/domain"
	"github.com/katzien/go-structure-examples/new/lib"
	"time"
)

// GetBeers returns all beers
// TODO - confirm if this should return pointers
func GetBeers() ([]domain.Beer, error) {
	return db.GetBeers()
}

// GetBeer returns a single beer
func GetBeer(id string) (*domain.Beer, error) {
	b, err := db.GetBeer(id)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// AddBeer persists the given beer to storage
// TODO - should this be in the domain??
// TODO - confirm pointer?
func AddBeer(b domain.Beer) (*domain.Beer, error) {
	// make sure we don't add any duplicates
	existingBeers, err := db.GetBeers()
	if err != nil {
		return nil, err
	}

	for _, e := range existingBeers {
		if b.Abv == e.Abv &&
			b.Brewery == e.Brewery &&
			b.Name == e.Name {
			return nil, domain.ErrDuplicateBeer
		}
	}

	// TODO: any other validation (e.g. valid brewery/ABV etc.) can be done here using calls to the domain

	id, err := lib.GetID("beer") // TODO should be a lib really, or domain func?
	if err != nil {
		return nil, err
	}
	b.ID = id

	b.Created = time.Now().UTC()

	err = db.AddBeer(b)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
