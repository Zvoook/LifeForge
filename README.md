# LifeForge

LifeForge is a learning Go project for managing personal development tasks and gradually building a real backend system around them.

The project started as a console task manager with JSON persistence, but is now evolving into a backend learning playground. It is used to practice Go, clean project structure, service/repository architecture, HTTP API development, JSON handling, error handling, testing, and later SQL/PostgreSQL.

## Current version

### v0.4 — CLI + basic HTTP API

Implemented features:

- create tasks
- show all tasks
- show tasks by area
- show tasks by status
- find task by ID
- edit task title
- edit task area
- edit task status
- edit task priority
- edit estimated minutes
- complete tasks
- delete tasks
- clear all tasks
- show dashboard
- build daily task plan
- save tasks to JSON
- load tasks from JSON on startup
- restore task IDs after loading saved data
- structured CLI output
- action preview before editing/deleting/completing tasks
- before/after output for task updates
- repository tests
- service tests
- storage tests
- validation tests
- area/status tests
- basic HTTP API server
- JSON HTTP responses
- centralized HTTP error responses
- production-like Go project structure

## Development areas

Tasks can belong to one of the following areas:

- Backend
- English
- Guitar
- Algorithms
- University

These areas reflect the main personal development directions tracked by the project.

## Project structure

```text
LifeForge/
├── cmd/
│   ├── lifeforge-cli/
│   │   └── main.go
│   │
│   └── lifeforge-api/
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
│   ├── cli/
│   │   ├── cli.go
│   │   ├── handlers.go
│   │   ├── readers.go
│   │   └── formatters.go
│   │
│   └── httpapi/
│       ├── server.go
│       ├── handlers.go
│       ├── requests.go
│       └── responses.go
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

The project is split into several layers.

1. cmd/lifeforge-cli

Console application entry point.

Responsibilities:

- load tasks from JSON
- create repository
- create task service
- create CLI application
- start terminal interface

Run:
```bush
go run ./cmd/lifeforge-cli
```

2. cmd/lifeforge-api

HTTP API application entry point.

Responsibilities:

- load tasks from JSON
- create repository
- create task service
- create HTTP API server
- start HTTP server on localhost:8080

Run:
```bush
go run ./cmd/lifeforge-api
```

3. internal/task

Domain and business logic layer.

Responsibilities:

- task entity
- development areas
- task statuses
- validation
- repository interface
- in-memory repository
- task service
- JSON storage
- business rules
- unit tests

The task package does not know anything about CLI or HTTP. It contains the core application logic.

4. internal/cli

Terminal user interface layer.

Responsibilities:

- menu rendering
- user input reading
- command handlers
- formatted task tables
- dashboard output
- daily plan output
- action preview before task modification
- saving changes after task operations

The CLI layer calls TaskService and prints human-readable output.

5. internal/httpapi

HTTP API layer.

Responsibilities:

- HTTP routing
- request handling
- JSON request decoding
- JSON response encoding
- HTTP status codes
- HTTP error responses
- mapping service errors to HTTP responses

The HTTP layer calls TaskService and returns machine-readable JSON responses.

## HTTP API

Base URL:
```url
http://localhost:8080
```

### 1. Health check
GET /health

Example response:
```shell
{
  "status": "ok"
}
```

### 2. Get all tasks
GET /tasks

Example response:
```shell
[
  {
    "id": 1,
    "area": 1,
    "title": "Learn HTTP in Go",
    "status": 1,
    "priority": 8,
    "estimatedMinutes": 60
  }
]
```

### 3. Create task
POST /tasks

Example request:
```shell
{
  "title": "Learn HTTP in Go",
  "area": 1,
  "priority": 8,
  "estimatedMinutes": 60
}
```
Example response:
```shell
{
  "id": 1,
  "area": 1,
  "title": "Learn HTTP in Go",
  "status": 1,
  "priority": 8,
  "estimatedMinutes": 60
}
```

### 4. Get task by ID
GET /tasks/{id}

Example:

GET /tasks/1

### 5. Delete task
DELETE /tasks/{id}

Example:

DELETE /tasks/1

The endpoint returns the deleted task.

### 6. Complete task
PATCH /tasks/{id}/complete

Example:

PATCH /tasks/1/complete

The endpoint marks the task as completed and returns the updated task.

## HTTP status codes

The API uses standard HTTP status codes:

```text
Status | Meaning
200    | OK	Successful request
201    | Created	Task created
400    | Bad Request	Invalid input, invalid ID, bad JSON
404    | Not Found	Task not found
405    | Method Not Allowed	HTTP method is not supported for this endpoint
409    | Conflict	Task state conflicts with the requested action
500    | Internal Server Error	Unexpected server-side error
```

Error response format:
```shell
{
  "error": "error message"
}
```

## JSON storage

Tasks are currently stored locally in:
```text
save.json
```
The file is ignored by Git because it contains local user data.

This storage approach is temporary. The next major development step is replacing JSON persistence with PostgreSQL.

## Run CLI
```bush
go run ./cmd/lifeforge-cli
```

## Run HTTP API
```bush
go run ./cmd/lifeforge-api
```

Then open:
```url
http://localhost:8080/health
```
or test with PowerShell:
```powershell
Invoke-RestMethod http://localhost:8080/health
```
### Build CLI
```bush
go build -o builds/lifeforge-task-cli.exe ./cmd/lifeforge-cli
```

### Run built executable on Windows PowerShell
```powershell
.\builds\lifeforge-task-cli.exe
```

## Tests

Run all tests:
```bush
go test ./...
```

Run only task package tests:
```bush
go test ./internal/task
```
## PowerShell ANSI colors

If colors are displayed incorrectly in PowerShell, run:
```powershell
reg add HKCU\Console /v VirtualTerminalLevel /t REG_DWORD /d 1
```

Then restart the terminal.

## Learning goals

This project is used for learning:

- Go basics
- structs and methods
- interfaces
- error handling
- package structure
- repository/service architecture
- CLI input/output
- JSON persistence
- HTTP servers
- HTTP routing
- REST-like API design
- JSON request/response handling
- HTTP status codes
- unit testing
- Git workflow
- GitHub Actions
- Roadmap

## Planned improvements:

- clean up HTTP API handlers
- add HTTP handler tests with httptest
- add PostgreSQL storage
- add SQL migrations
- add Docker Compose for API + PostgreSQL
- add configuration through environment variables
- add graceful shutdown
- add better dashboard API endpoint
- add daily plan API endpoint
- add task due dates
- add recurring tasks
- add study sessions
- add progress analytics
- add more tests for CLI-independent logic

## Long-term idea

LifeForge is intended to become a personal development backend system.

The long-term goal is to track tasks, learning areas, study sessions, daily plans, and progress analytics while using the project as a practical backend learning path.
