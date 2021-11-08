package storage

import (
	"encoding/json"
	"path"
	"runtime"
	"time"

	"github.com/katzien/go-structure-examples/new/domain"
	"github.com/nanobox-io/golang-scribble"
)

const (
	// dir defines the name of the directory where the files are stored
	dir = "/data/"

	// CollectionBeer identifier for the JSON collection of beers
	CollectionBeer = "beers"
	// CollectionReview identifier for the JSON collection of reviews
	CollectionReview = "reviews"
)

// Beer defines the storage form of a beer
type Beer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Brewery   string    `json:"brewery"`
	Abv       float32   `json:"abv"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}

// Review defines the storage form of a beer review
type Review struct {
	ID        string    `json:"id"`
	BeerID    string    `json:"beer_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	Added     time.Time `json:"added"`
}

// File storage keeps data in JSON files
type file struct {
	db *scribble.Driver
}

// NewFileStorage returns a new file storage
func NewFileStorage() (*file, error) {
	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	driver, err := scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return &file{db: driver}, nil
}

// AddBeer saves the given beer to the repository
func (s *file) AddBeer(b domain.Beer) error {
	bJSON := Beer{
		ID:        b.ID,
		Name:      b.Name,
		Brewery:   b.Brewery,
		Abv:       b.Abv,
		ShortDesc: b.ShortDesc,
		Created:   b.Created,
	}

	if err := s.db.Write(CollectionBeer, bJSON.ID, bJSON); err != nil {
		return err
	}

	return nil
}

// GetBeer returns a beer with the specified ID
func (s *file) GetBeer(ID string) (domain.Beer, error) {
	var b Beer

	if err := s.db.Read(CollectionBeer, ID, &b); err != nil {
		return domain.Beer{}, err
	}

	beer := domain.Beer{
		ID:        b.ID,
		Name:      b.Name,
		Brewery:   b.Brewery,
		Abv:       b.Abv,
		ShortDesc: b.ShortDesc,
		Created:   b.Created,
	}

	return beer, nil
}

// GetBeers returns all beers
func (s *file) GetBeers() ([]domain.Beer, error) {
	var list []domain.Beer

	records, err := s.db.ReadAll(CollectionBeer)
	if err != nil {
		return list, err
	}

	for _, r := range records {
		var b Beer
		if err := json.Unmarshal([]byte(r), &b); err != nil {
			return list, err
		}

		beer := domain.Beer{
			ID:        b.ID,
			Name:      b.Name,
			Brewery:   b.Brewery,
			Abv:       b.Abv,
			ShortDesc: b.ShortDesc,
			Created:   b.Created,
		}

		list = append(list, beer)
	}

	return list, nil
}

// AddReview saves the given review in the repository
func (s *file) AddReview(r domain.Review) error {
	rJSON := Review{
		ID:        r.ID,
		BeerID:    r.BeerID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
		Added:     r.Added,
	}

	if err := s.db.Write(CollectionReview, rJSON.ID, rJSON); err != nil {
		return err
	}

	return nil
}

// GetReviews returns all reviews for a given beer
func (s *file) GetReviews(beerID string) ([]domain.Review, error) {
	var list []domain.Review

	records, err := s.db.ReadAll(CollectionReview)
	if err != nil {
		return list, err
	}

	for _, b := range records {
		var r Review

		if err := json.Unmarshal([]byte(b), &r); err != nil {
			return list, err
		}

		if r.BeerID == beerID {
			review := domain.Review{
				ID:        r.ID,
				BeerID:    r.BeerID,
				FirstName: r.FirstName,
				LastName:  r.LastName,
				Score:     r.Score,
				Text:      r.Text,
				Added:     r.Added,
			}

			list = append(list, review)
		}
	}

	return list, nil
}
