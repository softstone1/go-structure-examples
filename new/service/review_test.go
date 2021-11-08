package service_test

import (
	"github.com/katzien/go-structure-examples/new/domain"
	"github.com/katzien/go-structure-examples/new/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReviewMustBeForExistingBeer(t *testing.T) {
	r := domain.Review{
		BeerID:    "does-not-exist-yet",
		FirstName: "Foo",
		LastName:  "Bar",
		Score:     5,
		Text:      "This is good stuff!",
	}

	_, err := service.AddBeerReview(r)
	require.Error(t, err)
	assert.EqualError(t, domain.ErrBeerNotFound, err.Error(), "expected an error because we're trying to add a review for a beer which doesn't exist")
}

func TestReviewService(t *testing.T) {
	b := domain.Beer{
		Name:      "Test Beer 1",
		Brewery:   "Brewery One",
		Abv:       3.6,
		ShortDesc: "Lorem Ipsum",
	}

	addedBeer, err := service.AddBeer(b)
	require.NoError(t, err)

	r := domain.Review{
		BeerID:    addedBeer.ID,
		FirstName: "Foo",
		LastName:  "Bar",
		Score:     5,
		Text:      "This is good stuff!",
	}

	addedReview, err := service.AddBeerReview(r)
	require.NoError(t, err)

	reviews, err := service.GetBeerReviews(addedBeer.ID)
	require.NoError(t, err)
	assert.Len(t, reviews, 1)
	assert.Equal(t, addedReview, reviews[0])
}
