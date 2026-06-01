package task

import (
	"fmt"
)

type Task struct {
	ID               uint32 `json:"id"`
	Area             Area   `json:"area"`
	Title            string `json:"title"`
	Status           Status `json:"status"`
	Priority         int    `json:"priority"`
	EstimatedMinutes int    `json:"estimatedMinutes"`
}

func (t *Task) Complete() {
	t.Status = Done
}

func (t *Task) ChangePriority(priority int) {
	t.Priority = priority
}

func (t Task) String() string {
	str := fmt.Sprintf("[%d] %s | %s | %v | priority: %d | %d min", t.ID, t.Area.String(), t.Title, t.Status.String(), t.Priority, t.EstimatedMinutes)
	return str
}
