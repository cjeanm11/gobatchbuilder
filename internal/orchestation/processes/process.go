package processes

import (
	"fmt"
	"gobatchbuilder/internal/orchestation/build"
)

func NewExBatchProcess() build.BatchProcess {
	return build.BatchProcess{
		Name: "ex_batch_porcess",
		Jobs: []build.Job{
			{
				Name: "Job1",
				Steps: build.Flow{
					Steps: []build.Step{
						{Name: "Step1", Executor: ExecuteStep1},
						{Name: "Step2", Executor: ExecuteStep2},
					},
				},
			},
		},
	}
}

func ExecuteStep1(_ any) (any, error) {
	fmt.Println("Executing Step1")
	var startingValue int8 = 1
	return any(startingValue), nil
}

func ExecuteStep2(input any) (any, error) {
	fmt.Println("Executing Step2")
	res := input.(int8)
	res += 1
	return any(res), nil
}
