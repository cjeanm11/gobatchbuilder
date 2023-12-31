package main

import (
	"fmt"
	"gobatch/config"
	"log"
)

func main() {
	// Default structure: Paths to config folder. Adjust according to your project.
	configPaths := []string{"../../config", "."}

	config, err := config.LoadConfig(configPaths)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Printf("Loaded Configuration: %+v\n", config)
}
