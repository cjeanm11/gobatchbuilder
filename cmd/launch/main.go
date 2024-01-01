package main

import (
	"fmt"
	"gobatchbuilder/config"
	"gobatchbuilder/internal/orchestation/batch"
	"gobatchbuilder/internal/orchestation/processes"
)

func main() {
	configPaths := []string{"../../config", "."}

	_, err := config.LoadConfig(configPaths)
	if err != nil {
		fmt.Printf("Error loading configurations: %v\n", err)
		return
	}

	// Create a spefific BatchProcess with jobs and steps
	batchProcess := processes.NewExBatchProcess()
	// Create batch process orchestrateur that will execute the process
	orchestrateur := batch.NewOrchestrator()

	if err := orchestrateur.ExecuteProcess(batchProcess); err != nil {
		fmt.Printf("Error executing process: %v\n", err)
		return
	}

	fmt.Println("Batch process completed successfully")
}
