package gobatchbuilder

import "fmt"

type Job struct {
	Name  string
	Steps Sequence
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

	sequencer := NewSequencer()

	for _, job := range process.Jobs {
		fmt.Printf("Executing job: %s\n", job.Name)

		finalState, err := sequencer.ExecuteSequence(&job.Steps)
		if err != nil {
			fmt.Printf("Error executing job %s: %s\n", job.Name, err)
			return err
		}

		fmt.Printf("Job %s completed successfully with final state: %+v\n", job.Name, finalState)
	}

	fmt.Printf("Process %s completed successfully\n", process.Name)
	return nil
}

type Sequence struct {
	Steps []Step
	State any
}

type Step interface {
	Execute(data interface{}) (interface{}, error)
}

func NewStep(name string, exec func(interface{}) (interface{}, error)) Step {
	return &genericStep{name: name, execute: exec}
}

type genericStep struct {
	name    string
	execute func(interface{}) (interface{}, error)
}

func (g *genericStep) Execute(data interface{}) (interface{}, error) {
	fmt.Printf("Executing %s\n", g.name)
	return g.execute(data)
}

type Sequencer struct{}

func NewSequencer() Sequencer {
	return Sequencer{}
}

func (s Sequencer) ExecuteSequence(sequence *Sequence) (interface{}, error) {
	currentState := sequence.State

	for _, step := range sequence.Steps {
		var err error
		currentState, err = step.Execute(currentState)
		if err != nil {
			return nil, fmt.Errorf("error executing step: %w", err)
		}
	}

	return currentState, nil
}
