package task

import "errors"

var (
	ErrInvalidId               = errors.New("ID must be more than 0")
	ErrInvalidTitle            = errors.New("task title cannot be empty")
	ErrInvalidArea             = errors.New("unknown area")
	ErrInvalidStatus           = errors.New("invalid task status")
	ErrInvalidPriority         = errors.New("priority must be between 1 and 10")
	ErrInvalidParameter        = errors.New("unknown parameter")
	ErrInvalidEstimatedMinutes = errors.New("estimated minutes must be greater than 0")
	ErrUnknownValue            = errors.New("unknown value")
	ErrTaskNotFound            = errors.New("task not found")
	ErrTaskAlreadyCompleted    = errors.New("task is already completed")
	ErrUnknownAction           = errors.New("unknown action")
)
