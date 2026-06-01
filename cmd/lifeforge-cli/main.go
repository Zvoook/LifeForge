package main

import (
	"fmt"

	"github.com/Zvoook/lifeforge/internal/cli"
	"github.com/Zvoook/lifeforge/internal/task"
)

const saveFileName = "save.json"

func main() {
	tasks, err := task.LoadTasksFromFile(saveFileName)
	if err != nil {
		fmt.Println("failed to load tasks:", err)
		return
	}

	repository := task.NewRepositoryFromTasks(tasks)
	service := task.NewTaskService(&repository)
	app := cli.NewCLI(&service, saveFileName)

	app.Run()
}
