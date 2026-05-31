package main

import "fmt"

func (c *CLI) handleCreateTask() {
	title, err := c.readTitle()
	if err != nil {
		printError(err)
		return
	}

	area, err := c.readArea()
	if err != nil {
		printError(err)
		return
	}

	priority, err := c.readPriority()
	if err != nil {
		printError(err)
		return
	}

	em, err := c.readEstimatedMinutes()
	if err != nil {
		printError(err)
		return
	}

	task, err := c.Service.CreateTask(title, area, priority, int(em))
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("\nCreated: \n%v\n", task)
}

func (c *CLI) handleShowAllTasks() {
	tasks := c.Service.GetAllTasks()

	for _, task := range tasks {
		fmt.Println(task)
	}
	fmt.Print("\n")
}

func (c *CLI) handleShowTasksByArea() {
	area, err := c.readArea()
	if err != nil {
		printError(err)
		return
	}

	tasks, err := c.Service.GetTasksByArea(area)
	if err != nil {
		printError(err)
		return
	}

	for _, task := range tasks {
		fmt.Println(task)
	}
	fmt.Print("\n")
}

func (c *CLI) handleFindTaskByID() {
	id, err := c.readID()
	if err != nil {
		printError(err)
		return
	}

	task, err := c.Service.GetTaskById(id)
	if err != nil {
		printError(err)
		return
	}

	fmt.Println(task)
	fmt.Print("\n")
}

func (c *CLI) handleCompleteTask() {
	id, err := c.readID()
	if err != nil {
		printError(err)
		return
	}

	_, err = c.Service.GetTaskById(id)
	if err != nil {
		printError(err)
		return
	}

	err = c.Service.CompleteTask(id)
	if err != nil {
		printError(err)
		return
	}

	printSuccess("Task completed")
}

func (c *CLI) handleChangeTaskPriority() {
	id, err := c.readID()
	if err != nil {
		printError(err)
		return
	}

	task, err := c.Service.GetTaskById(uint32(id))
	if err != nil {
		printError(err)
		return
	}
	fmt.Printf("Task #%d has priority %d\n", id, task.Priority)

	p, err := c.readPriority()
	if err != nil {
		printError(err)
		return
	}

	err = c.Service.ChangePriority(id, p)
	if err != nil {
		printError(err)
		return
	}

	printSuccess("Priority updated")
}

func (c *CLI) handleDeleteTask() {
	id, err := c.readID()
	if err != nil {
		printError(err)
		return
	}

	err = c.Service.DeleteTask(id)
	if err != nil {
		printError(err)
		return
	}

	printSuccess("Task deleted")
}

func (c *CLI) showDashboard() {
	printInfo("Not implemented yet")
}
