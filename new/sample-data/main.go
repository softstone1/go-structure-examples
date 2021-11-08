package main

import (
	"fmt"
	"github.com/katzien/go-structure-examples/new/service"
	"github.com/katzien/go-structure-examples/new/storage"
)

var (
	DB  service.Repository
	err error
)

func main() {
	DB, err = storage.NewFileStorage()
	if err != nil {
		fmt.Printf("😵 Failed to initialise file storage: %s\n", err.Error())
	}

	// add some sample beers
	for _, b := range DefaultBeers {
		_, err := service.AddBeer(b)
		if err != nil {
			fmt.Printf("😱 Error adding beer: %s\n", err.Error())
		}
	}

	fmt.Printf("\n✅ Added %d beers\n", len(DefaultBeers))

	// add some sample reviews
	for _, r := range DefaultReviews {
		_, err := service.AddBeerReview(r)
		if err != nil {
			fmt.Printf("😱 Error adding review: %s\n", err.Error())
		}
	}

	fmt.Printf("\n✅ Added %d reviews\n", len(DefaultReviews))

	fmt.Printf("\n🚀 Finished adding sample data\n")
}
