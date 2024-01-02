# Go batch builder

This Go package, `gobatchbuilder`, provides a simple structure for defining and executing basic batch processes. 

## Installation

You can install this package using Go modules:

```shell
go get github.com/cjeanm11/gobatchbuilder
```
## Quick Start
Here's a simple example to demonstrate how you might set up and execute a batch process:

##### 1. Implement a sequence of executable task of type `Step` interface: 

```go
    package main

    import (
        batch "github.com/cjeanm11/gobatchbuilder"
    )

    // Define steps for your batch process
    func FileReaderStep(data interface{}) (interface{}, error) {
        // Implement your file reading logic here...
    }

    func CapitalizeProcessorStep(data interface{}) (interface{}, error) {
        // Implement your data processing logic here...
    }

    func DatabaseWriterStep(data interface{}) (interface{}, error) {
        // Implement your database writing logic here...
    }
    
    // AnotherTask...
```
##### 2. Implement the tree-like structure where the root node represents the entire `BatchProcess`, and each child node represents a`Job` that needs to be executed. Each job can, in turn, have a `Sequence` of `Step`, which are the actual tasks that get executed.
```go
    func NewBatchProcess() batch.BatchProcess {
        // Define the root batch process with a descriptive name.
        return batch.BatchProcess{
            Name: "root",
            Jobs: []batch.Job{
                {
                    Name: "Job1", 
                    // List the steps to be executed as part of this job.
                    Sequence: batch.Sequence{
                        Steps: []batch.Step{
                            batch.NewStep("Read File", FileReaderStep),
                            batch.NewStep("Capitalize Data", CapitalizeProcessorStep),
                            batch.NewStep("Write to Database", DatabaseWriterStep),
                        },
                    },
                },
                // Add more jobs as needed.
            },
        }
    }
```
##### 3. Establishes an orchestrator responsible of managing the predefined jobs, and directs it to execute the process.
```go
    func main() {
        // Create an orchestrator to manage the batch process execution.
        orchestrateur := batch.NewOrchestrator() 
        // Initialize a new batch process.
        batchProcess := NewBatchProcess()

        // Execute the batch process and handle any errors.
        if err := orchestrateur.ExecuteProcess(batchProcess); err != nil {
            fmt.Printf("Error executing process: %v\n", err)
            return
        }

        fmt.Println("Batch process completed successfully")
    }
```