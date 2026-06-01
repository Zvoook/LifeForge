package main

import (
	"github.com/Zvoook/lifeforge/internal/cli"
	"github.com/Zvoook/lifeforge/internal/task"
)

const saveFileName = "save.json"

func main() {
	repository, err := task.NewRepositoryFromFile(saveFileName)
	if err != nil {
		cli.PrintError(err)
		return
	}
	service := task.NewTaskService(&repository)
	app := cli.NewCLI(&service, saveFileName)

	app.Run()
}
