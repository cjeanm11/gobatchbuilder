package orchestrator

import (
	"fmt"
	"gobatchbuilder/internal/orchestation/sequencer"
)

type BatchJob struct {
	Name  string
	Steps []sequencer.Step
}

type Orchestrator struct{}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{}
}

func (o *Orchestrator) OrchestrateJob(job BatchJob) error {
	seq := sequencer.NewSequencer()
	fmt.Printf("Orchestrating job: %s\n", job.Name)
	for _, step := range job.Steps {
		fmt.Printf("Executing step: %s\n", step.Name)
		if err := seq.ExecuteStep(step); err != nil {
			return err
		}
	}
	fmt.Printf("Job %s completed successfully\n", job.Name)
	return nil
}
