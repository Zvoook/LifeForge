package main

func main() {
	tasks, err := LoadTasksFromFile("save.json")

	if err != nil {
		printError(err)
		return
	}

	repository := NewRepositoryFromTasks(tasks)
	service := NewTaskService(&repository)
	cli := NewCLI(&service)

	cli.run()
}
