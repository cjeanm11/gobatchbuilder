package main

import (
	"fmt"
	"gobatchbuilder/config"
	"gobatchbuilder/internal/orchestation/orchestrator"
	"gobatchbuilder/internal/orchestation/sequencer"
	"log"
)

func main() {
	configPaths := []string{"../../config", "."}

	configuration, err := config.LoadConfig(configPaths)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Printf("Loaded Configuration: %+v\n", configuration)
	o := orchestrator.NewOrchestrator()
	job := orchestrator.BatchJob{
		Name: "ExampleJob",
		Steps: []sequencer.Step{
			{
				Name: "Step1",
				Executor: func() error {
					fmt.Println("Executing Step1")
					return nil
				},
			},
			{
				Name: "Step2",
				Executor: func() error {
					fmt.Println("Executing Step2")
					return nil
				},
			},
		},
	}

	if err := o.OrchestrateJob(job); err != nil {
		fmt.Printf("Job orchestration error: %v\n", err)
	}
}
