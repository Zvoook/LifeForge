package main

import (
	"fmt"
)

func main() {
	task1 := Task{
		id:               1,
		area:             Backend,
		title:            "First task",
		status:           in_progress,
		priority:         1,
		estimatedMinutes: 120,
	}

	fmt.Println(task1)
}
