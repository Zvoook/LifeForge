# LifeForge

LifeForge is a learning Go project for managing personal development tasks.

The current version is a console task manager.
It helps create, view, update, complete and delete tasks grouped by development areas.

## Current version
### v0.1.0 — LifeForge Task CLI
Implemented features:
- create tasks
- show all tasks
- show tasks by area
- find task by ID
- complete tasks
- change task priority
- delete tasks
- show basic dashboard
- in-memory task repository

## Development areas
Tasks can belong to one of the following areas:
- Backend
- English
- Guitar
- Algorithms
- University

## Run

```bash
go run ./playground/TaskService
```
## Build
```bash
go build -o builds/lifeforge-task-cli.exe ./playground/TaskService
```

## Run built executable
```bash
./builds/lifeforge-task-cli.exe
```

## PowerShell ANSI colors
If colors are displayed incorrectly in PowerShell, run:
```shell
reg add HKCU\Console /v VirtualTerminalLevel /t REG_DWORD /d 1
```
Then restart the terminal.

## Project status
This project is currently used for learning:
- Go basics
- structs and methods
- interfaces
- repository/service architecture
- CLI input/output
- error handling

## Planned improvements:
- save tasks to JSON file
- load tasks on application start
- add tests for service and repository layers
- move code from playground to a production-like project structure