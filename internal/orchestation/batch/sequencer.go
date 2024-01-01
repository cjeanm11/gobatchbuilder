package batch

import "fmt"

type Flow []Step

type Step struct {
	Name     string
	Executor func() error
}

type Sequencer struct{}

func NewSequencer() *Sequencer {
	return &Sequencer{}
}

func (s *Sequencer) executeStep(step Step) error {
	fmt.Printf("Executing step: %s\n", step.Name)
	if err := step.Executor(); err != nil {
		return err
	}
	fmt.Printf("Step %s completed successfully\n", step.Name)
	return nil
}

func (s *Sequencer) ExecuteFlow(flow Flow) error {
	for _, step := range flow {
		err := s.executeStep(step)
		if err != nil {
			return err
		}
	}
	return nil
}
