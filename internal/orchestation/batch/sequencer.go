package batch

import "fmt"

type Flow struct {
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

func (s Sequencer) ExecuteFlow(flow *Flow) (interface{}, error) {
	currentState := flow.State

	for _, step := range flow.Steps {
		var err error
		currentState, err = step.Execute(currentState)
		if err != nil {
			return nil, fmt.Errorf("error executing step: %w", err)
		}
	}

	return currentState, nil
}
