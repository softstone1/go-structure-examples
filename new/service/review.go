package service

import (
	"github.com/katzien/go-structure-examples/new/domain"
	"github.com/katzien/go-structure-examples/new/lib"
	"time"
)

// GetBeerReviews returns all requests for a beer
func GetBeerReviews(beerID string) ([]domain.Review, error) {
	return db.GetReviews(beerID)
}

// AddBeerReview saves a new beer review
func AddBeerReview(r domain.Review) (*domain.Review, error) {
	// make sure we're adding a review for an existing beer
	_, err := db.GetBeer(r.BeerID)
	if err != nil {
		return nil, err
	}

	id, err := lib.GetID("review") // TODO should be a lib really, or domain func?
	if err != nil {
		return nil, err
	}
	r.ID = id

	r.Added = time.Now().UTC()

	err = db.AddReview(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
