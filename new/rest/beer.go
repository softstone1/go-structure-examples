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

// Beer defines the properties of a beer
type Beer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Brewery   string    `json:"brewery"`
	Abv       float32   `json:"abv"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}

// getBeer returns a handler for GET /beers/:id requests
func getBeer() func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		beer, err := service.GetBeer(p.ByName("id"))
		if err != nil {
			if err == domain.ErrBeerNotFound {
				http.Error(w, "the beer you requested does not exist", http.StatusNotFound)
				return
			}

			http.Error(w, fmt.Sprintf("failed to get beer: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// TODO - marshall between domain and rest type?

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(beer)
	}
}

// getBeers returns a handler for GET /beers requests
func getBeers() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list, err := service.GetBeers()
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get beers: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// TODO - marshall between domain and rest type?

		json.NewEncoder(w).Encode(list)
	}
}

// addBeer returns a handler for POST /beers requests
func addBeer() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var b Beer
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&b)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse beer: %s", err.Error()), http.StatusBadRequest)
			return
		}

		// TODO - double check if we need the marshalling here?

		newBeer := domain.Beer{
			Name:      b.Name,
			Brewery:   b.Brewery,
			Abv:       b.Abv,
			ShortDesc: b.ShortDesc,
		}

		beer, err := service.AddBeer(newBeer)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to add beer: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(beer) // TODO - check how this gets returned, need to marshall?
	}
}
