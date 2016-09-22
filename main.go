package main

import (
	"github.com/BelfastGophers/gophurl/data"
	"github.com/BelfastGophers/gophurl/services"

	"fmt"
)

const (
	// Path to a JSON file
	pathToJsonFile = "demo.json" // Eeewwwwwww, code smell
)

// run is called by main so we can actually test this logic
// TODO: Write tests....
func run() error {
	repo, err := data.NewJsonShortURLRepository(pathToJsonFile)
	if err != nil {
		return err
	}

	svc := services.NewShortener(repo)
	return svc.Start(8000)
}

// Entrypoint function
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
