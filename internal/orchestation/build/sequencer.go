package build

import (
	"fmt"
)

type Flow struct {
	Steps  []Step
	Result interface{}
}

type Step struct {
	Name     string
	Executor func(any) (any, error)
}

type Sequencer struct{}

func NewSequencer() Sequencer {
	return Sequencer{}
}
func (s *Sequencer) ExecuteFlow(flow Flow) (interface{}, error) {
	for _, step := range flow.Steps {
		res, err := step.Executor(flow.Result)
		if err != nil {
			fmt.Printf("Error executing step: %s\n", err)
			return nil, err
		}
		flow.Result = res
		fmt.Printf("Result after %s: %v\n", step.Name, res)
	}
	return flow.Result, nil
}
