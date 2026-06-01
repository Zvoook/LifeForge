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