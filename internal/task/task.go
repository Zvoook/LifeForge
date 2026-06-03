package task

import (
	"fmt"
	"strings"
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

func (t *Task) EditPriority(priority int) {
	t.Priority = priority
}

func (t *Task) EditTitle(title string) {
	t.Title = strings.TrimSpace(title)
}

func (t *Task) EditArea(area Area) {
	t.Area = area
}

func (t *Task) EditStatus(status Status) {
	t.Status = status
}

func (t *Task) EditEstimatedMinutes(minutes int) {
	t.EstimatedMinutes = minutes
}

func (t Task) String() string {
	str := fmt.Sprintf("[%d] %s | %s | %v | priority: %d | %d min", t.ID, t.Area.String(), t.Title, t.Status.String(), t.Priority, t.EstimatedMinutes)
	return str
}

func IsDuplicates(task_1 Task, task_2 Task) bool {
	return task_1.Title == task_2.Title && task_1.Area == task_2.Area &&
		task_1.Status == task_2.Status && task_1.Priority == task_2.Priority &&
		task_1.EstimatedMinutes == task_2.EstimatedMinutes
}
