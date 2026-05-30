package main

import (
	"fmt"
)

func main() {
	rep := NewRepository()
	service := NewTaskService(&rep)

	service.CreateTask("first task", Backend, 1, 120)
	service.CreateTask("second task", Backend, 2, 120)
	service.CreateTask("third task", Backend, 3, 120)

	tasks := service.GetAllTasks()
	for _, task := range tasks {
		fmt.Println(task)
	}

	err := service.CompleteTask(1)
	fmt.Println(err)

	err = service.CompleteTask(1)
	fmt.Println(err)

	err = service.CreateTask("", English, 4, 120)
	fmt.Println(err)

	_, err = service.GetTaskById(999)
	fmt.Println(err)
}
