package main

import "errors"

var (
	ErrEmptyTaskTitle          = errors.New("task title cannot be empty")
	ErrUnknownArea             = errors.New("unknown area")
	ErrInvalidTaskStatus       = errors.New("invalid task status")
	ErrInvalidPriority         = errors.New("priority must be between 1 and 10")
	ErrInvalidEstimatedMinutes = errors.New("estimated minutes must be greater than 0")
	ErrTaskNotFound            = errors.New("task not found")
	ErrTaskAlreadyCompleted    = errors.New("task is already completed")
)
