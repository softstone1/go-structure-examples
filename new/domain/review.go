package domain

import "time"

// Review defines a beer review
type Review struct {
	ID        string
	BeerID    string
	FirstName string
	LastName  string
	Score     int
	Text      string
	Added     time.Time
}

type ReviewRepository interface {
	AddReview(r Review) error
	GetReviews(beerID string) ([]Review, error)
}
