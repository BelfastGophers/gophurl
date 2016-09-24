package main

import (
	"github.com/AndrewSpeed/gophurl/data"
	"github.com/AndrewSpeed/gophurl/services"

	"fmt"
)

const (
	// Path to a JSON file
	pathToJSONFile = "demo.json" // Eeewwwwwww, code smell
)

// run is called by main so we can actually test this logic
// TODO: Write tests....
func run() error {
	repo, err := data.NewJSONShortURLRepository(pathToJSONFile)
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
