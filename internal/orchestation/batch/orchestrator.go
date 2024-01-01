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
		s := NewSequencer()
		for _, step := range job.Steps {
			fmt.Printf("Executing step: %s\n", step.Name)
			if err := s.ExecuteStep(step); err != nil {
				return err
			}
		}

		fmt.Printf("Job %s completed successfully\n", job.Name)
	}

	fmt.Printf("Process %s completed successfully\n", process.Name)
	return nil
}
