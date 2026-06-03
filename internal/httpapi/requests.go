package httpapi

import (
	"github.com/Zvoook/lifeforge/internal/task"
)

type CreateTaskRequest struct {
	Title            string    `json:title`
	Area             task.Area `json:area`
	Priority         int       `json:priority`
	EstimatedMinutes int       `json:estimatedMinutes`
}
