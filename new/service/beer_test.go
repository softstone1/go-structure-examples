package service_test

import (
	"github.com/katzien/go-structure-examples/new/domain"
	"github.com/katzien/go-structure-examples/new/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBeerService(t *testing.T) {
	b := domain.Beer{
		Name:      "Test Beer 1",
		Brewery:   "Brewery One",
		Abv:       3.6,
		ShortDesc: "Lorem Ipsum",
	}

	addedBeer, err := service.AddBeer(b)
	require.NoError(t, err)

	beers, err := service.GetBeers()
	require.NoError(t, err)
	assert.Len(t, beers, 1)
	assert.Equal(t, addedBeer, beers[0])

	beer, err := service.GetBeer(addedBeer.ID)
	require.NoError(t, err)
	assert.Equal(t, addedBeer, beer)
}
