package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/katzien/go-structure-examples/new/domain"
	"github.com/katzien/go-structure-examples/new/service"
	"net/http"
	"time"
)

// Review defines a beer review
type Review struct {
	ID        string    `json:"id"`
	BeerID    string    `json:"beer_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}

// getBeerReviews returns a handler for GET /beers/:id/reviews requests
func getBeerReviews() func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		reviews, err := service.GetBeerReviews(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get reviews: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviews) // TODO - need marshalling?
	}
}

// addBeerReview returns a handler for POST /beers/:id/reviews requests
func addBeerReview() func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		var r Review
		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&r); err != nil {
			http.Error(w, fmt.Sprintf("failed to parse review: %s", err.Error()), http.StatusBadRequest)
		}

		newReview := domain.Review{
			BeerID:    p.ByName("id"),
			FirstName: r.FirstName,
			LastName:  r.LastName,
			Score:     r.Score,
			Text:      r.Text,
		}

		review, err := service.AddBeerReview(newReview)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to add review: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(review) // TODO marshalling?
	}
}
