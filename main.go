package main

import (
	"fmt"
	"gobatchbuilder/internal/orchestation/batch"
	"gobatchbuilder/internal/orchestation/processes"
)

func main() {

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
