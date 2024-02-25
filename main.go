package main

import (
	"log"

	"github.com/atmatm9182/pomegranate/blueprint"
)

const defaultConfigPath = "config.yaml"

func main() {
	b, err := blueprint.Parse(defaultConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	
	err = b.Scaffold()
	if err != nil {
		log.Fatal(err)
	}
}
