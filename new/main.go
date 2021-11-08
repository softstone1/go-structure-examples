package main

import (
	"fmt"
	http2 "github.com/katzien/go-structure-examples/new/rest"
	"github.com/katzien/go-structure-examples/new/service"
	"github.com/katzien/go-structure-examples/new/storage"
	"log"
	"net/http"
)

// Type defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
)

var (
	db  service.Repository
	err error
)

func main() {
	// set up storage
	storageType := Memory // this could be a flag; hardcoded here for simplicity

	switch storageType {
	case Memory:
		db = &storage.Memory{}

	case JSON:
		db, err = storage.NewFileStorage()
		if err != nil {
			log.Fatalf("failed to initialise file storage: %v", err)
			return
		}
	}

	service.SetDB(db)

	// set up the HTTP server
	router := http2.Handle()

	fmt.Println("🍻 Beer server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
