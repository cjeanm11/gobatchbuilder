package processes

import (
	"fmt"
	batch "github.com/cjeanm11/gobatchbuilder/internal/orchestation/batch"
	"strings"
)

func NewExBatchProcess() batch.BatchProcess {
	return batch.BatchProcess{
		Name: "ex_batch_process",
		Jobs: []batch.Job{
			{
				Name: "Job1",
				Steps: batch.Sequence{
					Steps: []batch.Step{
						batch.NewStep("Read File", FileReaderStep),
						batch.NewStep("Capitalize Data", CapitalizeProcessorStep),
						batch.NewStep("Write to Database", DatabaseWriterStep),
					},
				},
			},
		},
	}
}

func FileReaderStep(data interface{}) (interface{}, error) {
	fmt.Println("Step 1: Read File")
	filePath := "file_path"
	readData := []string{"data1 from " + filePath, "data2 from " + filePath}
	return readData, nil
}

func CapitalizeProcessorStep(data interface{}) (interface{}, error) {
	fmt.Println("Step 2: Capitalize Data")
	strData, ok := data.([]string)
	if !ok {
		return nil, fmt.Errorf("expected []string, got %T", data)
	}
	var processedData []string
	for _, str := range strData {
		processedData = append(processedData, strings.ToUpper(str))
	}

	fmt.Println("Step 2: Processed data: ", processedData)
	return processedData, nil
}

func DatabaseWriterStep(data interface{}) (interface{}, error) {
	fmt.Println("Step 3: Write to Database")
	// Placeholder for actual database writing logic
	// ...

	strData, ok := data.([]string)
	if !ok {
		return nil, fmt.Errorf("expected []string, got %T", data)
	}
	fmt.Println("Step 3: Written: ", strData)
	writeStatus := "Data written successfully"
	return writeStatus, nil
}
