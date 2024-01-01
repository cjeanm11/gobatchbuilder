package sequencer

import "fmt"

type Step struct {
	Name     string
	Executor func() error
}

type Sequencer struct{}

func NewSequencer() Sequencer {
	return Sequencer{}
}

func (s *Sequencer) ExecuteStep(step Step) error {
	fmt.Printf("Executing step: %s\n", step.Name)
	if err := step.Executor(); err != nil {
		return err
	}
	fmt.Printf("Step %s completed successfully\n", step.Name)
	return nil
}
