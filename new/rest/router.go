package rest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Handle() http.Handler {
	router := httprouter.New()

	router.GET("/beers", getBeers())
	router.GET("/beers/:id", getBeer())
	router.POST("/beers", addBeer())

	router.GET("/beers/:id/reviews", getBeerReviews())
	router.POST("/beers/:id/reviews", addBeerReview())

	return router
}
