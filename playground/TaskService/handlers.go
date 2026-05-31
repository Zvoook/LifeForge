package main

import "fmt"

func (c *CLI) printMenu() {
	fmt.Println(Menu)
}

func (c *CLI) printSuccess(prompt string) {
	fmt.Sprint("✓ %s", prompt)
}

func (c *CLI) printError(err error) {
	fmt.Sprint("✗ Error: %v", err)
}

func (c *CLI) printInfo(prompt string) {
	fmt.Sprint("[i] %s", prompt)
}

func (c *CLI) handleCreateTask() error {
	title, err := c.readTitle()
	if err != nil {
		return err
	}
	area, err := c.readArea()
	if err != nil {
		return err
	}
	priority, err := c.readPriority()
	if err != nil {
		return err
	}
	estMin, err := c.readEstimatedMinutes()
	if err != nil {
		return err
	}

	task, err := c.Service.CreateTask(title, area, priority, estMin)
	if err != nil {
		return err
	}
	fmt.Printf("Created: \n%v", task)
	return nil
}

func (c *CLI) handleShowAllTasks() {
	tasks := c.Service.GetAllTasks()

	for _, task := range tasks {
		fmt.Println(task)
	}
	fmt.Print("\n")
}

func (c *CLI) handleShowTasksByArea() error {
	area, err := c.readArea()
	if err != nil {
		return err
	}

	tasks, err := c.Service.GetTasksByArea(area)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		fmt.Println(task)
	}
	fmt.Print("\n")
	return nil
}

func (c *CLI) handleFindTaskByID() error {
	id, err := c.readID()
	if err != nil {
		return err
	}

	task, err := c.Service.GetTaskById(id)
	if err != nil {
		return err
	}

	fmt.Println(task)
	fmt.Println("\n")
	return nil
}

func (c *CLI) handleCompleteTask() error {
	id, err := c.readID()
	if err != nil {
		return err
	}

	_, err = c.Service.GetTaskById(id)
	if err != nil {
		return err
	}

	err = c.Service.CompleteTask(id)
	if err != nil {
		return err
	}

	c.printSuccess("Task completed")
	return nil
}

func (c *CLI) handleChangeTaskPriority() error {
	id, err := c.readID()
	if err != nil {
		return err
	}
	p, err := c.readPriority()
	if err != nil {
		return err
	}

	err = c.Service.ChangePriority(id, p)
	if err != nil {
		return err
	}

	c.printSuccess("Priority updated")
	return nil
}

func (c *CLI) handleDeleteTask() error {
	id, err := c.readID()
	if err != nil {
		return err
	}

	err = c.Service.DeleteTask(id)
	if err != nil {
		return err
	}

	c.printSuccess("Task deleted")
	return nil
}

func (c *CLI) showDashboard() {
	c.printInfo("Not implemented yet")
}
