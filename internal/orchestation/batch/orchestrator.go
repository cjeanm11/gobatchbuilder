package gobatchbuilder

import "fmt"

type Job struct {
	Name     string
	Sequence Sequence
}

type BatchProcess struct {
	Name string
	Jobs []Job
}

type Orchestrator struct{}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{}
}

func (o *Orchestrator) ExecuteProcess(process BatchProcess) error {
	fmt.Printf("Executing process: %s\n", process.Name)

	sequencer := NewSequencer() // Create a sequencer instance

	for _, job := range process.Jobs {
		fmt.Printf("Executing job: %s\n", job.Name)

		finalState, err := sequencer.ExecuteSequence(&job.Sequence)
		if err != nil {
			fmt.Printf("Error executing job %s: %s\n", job.Name, err)
			return err
		}

		fmt.Printf("Job %s completed successfully with final state: %+v\n", job.Name, finalState)
	}

	fmt.Printf("Process %s completed successfully\n", process.Name)
	return nil
}
