package service

import "github.com/katzien/go-structure-examples/new/domain"

var db Repository

func SetDB(r Repository) {
	db = r
}

type Repository interface {
	domain.BeerRepository
	domain.ReviewRepository
}
