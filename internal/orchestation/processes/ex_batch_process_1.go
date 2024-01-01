package processes

import (
	"fmt"
	"gobatchbuilder/internal/orchestation/batch"
)

func NewExBatchProcess() batch.BatchProcess {
	return batch.BatchProcess{
		Name: "ex_batch_porcess",
		Jobs: []batch.Job{
			{
				Name: "Job1",
				Steps: batch.Flow{
					{Name: "Step1", Executor: executeStep1},
					{Name: "Step2", Executor: executeStep2},
					// add more steps...
				},
			},
		},
	}
}

func executeStep1() error {
	fmt.Println("Executing Step1")
	return nil
}

func executeStep2() error {
	fmt.Println("Executing Step2")
	return nil
}
