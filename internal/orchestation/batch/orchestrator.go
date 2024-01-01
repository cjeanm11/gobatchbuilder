package batch

import (
	"fmt"
)

type BatchProcess struct {
	Name string
	Jobs []Job
}

type Job struct {
	Name  string
	Steps Flow
}

type Orchestrator struct{}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{}
}

func (o *Orchestrator) ExecuteProcess(process BatchProcess) error {
	fmt.Printf("Executing process: %s\n", process.Name)

	for _, job := range process.Jobs {
		fmt.Printf("Executing job: %s\n", job.Name)
		sequencer := NewSequencer()
		if err := sequencer.ExecuteFlow(job.Steps); err != nil {
			fmt.Printf("Error executing flow: %s\n", err)
			return err
		}
		fmt.Printf("Job %s completed successfully\n", job.Name)
	}

	fmt.Printf("Process %s completed successfully\n", process.Name)
	return nil
}
