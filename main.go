package main

import (
	"fmt"
	batch "github.com/cjeanm11/gobatchbuilder/internal/orchestation/batch"
	"github.com/cjeanm11/gobatchbuilder/internal/orchestation/processes"
)

func main() {

	batchProcess := processes.NewExBatchProcess()
	// Create batch process orchestrateur that will execute the process
	orchestrateur := batch.NewOrchestrator()

	if err := orchestrateur.ExecuteProcess(batchProcess); err != nil {
		fmt.Printf("Error executing process: %v\n", err)
		return
	}

}
