package main

import (
	"fmt"
	"log"
	"os"

	"github.com/atmatm9182/pomegranate/blueprint"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: expected at least one argument\n")
		os.Exit(1)
	}

	repo := args[0]
	b, err := blueprint.FromRepo(repo, "pomegranate.yml")
	if err != nil {
		log.Fatalln(err)
	}

	err = b.Scaffold()
	if err != nil {
		log.Fatalln(err)
	}
}
