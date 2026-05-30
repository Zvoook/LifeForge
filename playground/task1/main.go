package main

import (
	"fmt"
)

func main() {
	task1 := Task{
		ID:               1,
		Area:             Backend,
		Title:            "First task",
		Status:           in_progress,
		Priority:         1,
		EstimatedMinutes: 120,
	}

	fmt.Println(task1)
}
