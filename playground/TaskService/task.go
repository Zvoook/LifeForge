package main

import (
	"fmt"
)

type Task struct {
	ID               uint32
	Area             Area
	Title            string
	Status           Status
	Priority         int
	EstimatedMinutes int
}

func (t *Task) Complete() {
	t.Status = done
}

func (t *Task) ChangePriority(priority int) {
	t.Priority = priority
}

func (t Task) String() string {
	str := fmt.Sprintf("[%d] %s | %s | %v | priority: %d | %d min", t.ID, t.Area.String(), t.Title, t.Status.String(), t.Priority, t.EstimatedMinutes)
	return str
}
