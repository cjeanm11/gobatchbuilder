cmd/batchmain/main.go: The starting point of the application, configuring and launching batch jobs.

internal/: Core application logic that isn't exposed externally.

orchestrator/: Manages the execution flow of batch jobs and steps.
components/: Specific batch processing functionalities like reading, processing, and writing data.
common/: Shared utilities and helpers such as logging and validation.
pkg/: Libraries meant to be used by external applications, containing configuration loading and shared models.