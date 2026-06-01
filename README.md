# LifeForge
LifeForge is a learning Go project for managing personal development tasks from the terminal.

The current version is a console task manager with JSON-based persistence. It helps create, view, update, complete, delete and analyze tasks grouped by development areas.

## Current version
### v0.3 — Structured CLI project with tests

Implemented features:
- create tasks
- show all tasks
- show tasks by area
- find task by ID
- complete tasks
- change task priority
- delete tasks
- clear all tasks
- show basic dashboard
- save tasks to JSON
- load tasks from JSON on startup
- restore task IDs after loading saved data
- repository tests
- service tests
- storage tests
- validation tests
- area/status tests
- production-like Go project structure

## Development areas
Tasks can belong to one of the following areas:
- Backend
- English
- Guitar
- Algorithms
- University

## Project structure
```text
LifeForge/
├── cmd/
│   └── lifeforge-cli/
│       └── main.go
│
├── internal/
│   ├── task/
│   │   ├── area.go
│   │   ├── status.go
│   │   ├── task.go
│   │   ├── errors.go
│   │   ├── validation.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   ├── storage.go
│   │   └── *_test.go
│   │
│   └── cli/
│       ├── cli.go
│       ├── handlers.go
│       ├── readers.go
│       └── formatters.go
│
├── .github/
│   └── workflows/
│       └── go-tests.yml
│
├── .gitignore
├── go.mod
└── README.md
```

## Architecture

The project is split into three main parts:

1) cmd/lifeforge-cli
Application entry point.

Responsibilities:
- load tasks from JSON
- create repository
- create service
- create CLI
- start the application


2) internal/task
Task domain and business logic.

Responsibilities:
- task entity
- development areas
- task statuses
- validation
- repository interface
- in-memory repository
- task service
- JSON storage
- unit tests


3) internal/cli
Console user interface.

Responsibilities:
- menu rendering
- user input reading
- command handlers
- formatted output
- saving changes after task operations

## Run
```bash
go run ./cmd/lifeforge-cli
```
## Build
```bash
go build -o builds/lifeforge-task-cli.exe ./cmd/lifeforge-cli
```

## Run built executable
```bash
./builds/lifeforge-task-cli.exe
```

### On Windows PowerShell:
```powershell
.\builds\lifeforge-task-cli.exe
```

## Tests

### Run all tests:

```bash
go test ./...
```

### Run only task package tests:

```bash
go test ./internal/task
```

## JSON storage

Tasks are stored locally in "save.json"

The file is ignored by Git because it contains local user data.

## PowerShell ANSI colors

If colors are displayed incorrectly in PowerShell, run:

```powershell
reg add HKCU\Console /v VirtualTerminalLevel /t REG_DWORD /d 1
```

Then restart the terminal.

## Project status

This project is currently used for learning:
- Go basics
- structs and methods
- interfaces
- repository/service architecture
- package structure
- CLI input/output
- error handling
- JSON persistence
- unit testing
- Git workflow
- GitHub Actions
- Planned improvements
- improve GitHub Actions workflow
- add release builds
- add better dashboard statistics
- improve CLI UX
- add task filtering by status and priority
- add task editing
- add due dates
- add recurring tasks
- add more tests for CLI-independent logic