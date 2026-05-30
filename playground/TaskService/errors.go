package main

import "errors"

var (
	ErrInvalidTitle            = errors.New("task title cannot be empty")
	ErrInvalidArea             = errors.New("unknown area")
	ErrInvalidStatus           = errors.New("invalid task status")
	ErrInvalidPriority         = errors.New("priority must be between 1 and 10")
	ErrInvalidEstimatedMinutes = errors.New("estimated minutes must be greater than 0")
	ErrTaskNotFound            = errors.New("task not found")
	ErrTaskAlreadyCompleted    = errors.New("task is already completed")
)
