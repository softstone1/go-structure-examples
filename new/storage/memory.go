package storage

import (
	"github.com/katzien/go-structure-examples/new/domain"
)

// Memory storage keeps data in memory
type Memory struct {
	beers   []domain.Beer
	reviews map[string][]domain.Review // map indexed by beer ID
}

// GetBeer returns a beer with the specified ID
func (m *Memory) GetBeer(ID string) (domain.Beer, error) {
	for _, b := range m.beers {
		if b.ID == ID {
			return b, nil
		}
	}

	return domain.Beer{}, domain.ErrBeerNotFound
}

// GetBeers returns all beers
func (m *Memory) GetBeers() ([]domain.Beer, error) {
	return m.beers, nil
}

// AddBeer saves the given beer to the in-memory database
func (m *Memory) AddBeer(b domain.Beer) error {
	m.beers = append(m.beers, b)
	return nil
}

// AddReview saves the given review in the in-memory database
func (m *Memory) AddReview(r domain.Review) error {
	list, ok := m.reviews[r.BeerID]
	if !ok {
		// we haven't had a review for this beer yet, so we add a new map index
		m.reviews[r.BeerID] = []domain.Review{r}
		return nil
	}

	// we already have other reviews for this beer, so we add it to the list
	m.reviews[r.BeerID] = append(list, r)
	return nil
}

// GetReviews returns all reviews for a given beer
func (m *Memory) GetReviews(beerID string) ([]domain.Review, error) {
	return m.reviews[beerID], nil
}
