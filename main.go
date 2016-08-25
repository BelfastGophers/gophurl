package main

import (
	"flag"
	"fmt"
)

var jsonRepoPath string

func init() {
	flag.StringVar(&jsonRepoPath, "json", "", "Path to JSON repository")
	flag.Parse()
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
