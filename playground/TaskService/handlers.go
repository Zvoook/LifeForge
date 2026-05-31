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
	tasks := c.Service.GetAllTasks()
	var todo_cnt, completed_cnt, blocked_cnt, cancelled_cnt, in_progress_cnt int
	var totalTimeCount, highPriorityTaskCount int
	var backend_cnt, english_cnt, algorithms_cnt, guitar_cnt, university_cnt int

	for i := 0; i < len(tasks); i++ {
		totalTimeCount += tasks[i].EstimatedMinutes

		if tasks[i].Priority >= 8 {
			highPriorityTaskCount++
		}

		switch tasks[i].Status {
		case todo:
			todo_cnt++
		case done:
			completed_cnt++
		case blocked:
			blocked_cnt++
		case cancelled:
			cancelled_cnt++
		case in_progress:
			in_progress_cnt++
		}

		switch tasks[i].Area {
		case Backend:
			backend_cnt++
		case English:
			english_cnt++
		case Algorithms:
			algorithms_cnt++
		case Guitar:
			guitar_cnt++
		case University:
			university_cnt++
		}
	}

	fmt.Printf("%sLifeForge Dashboard%s\n\n", Yellow, ResetColor)

	fmt.Printf("Total tasks: %d\n", len(tasks))
	fmt.Printf("Completed: %d\n", completed_cnt)
	fmt.Printf("To Do: %d\n", todo_cnt)
	fmt.Printf("In Progress: %d\n", in_progress_cnt)
	fmt.Printf("Blocked: %d\n", blocked_cnt)
	fmt.Printf("Cancelled: %d\n", cancelled_cnt)
	fmt.Print("\n")

	fmt.Printf("Total estimated time: %d\n", totalTimeCount)
	fmt.Printf("High priority tasks: %d\n", highPriorityTaskCount)
	fmt.Print("\n")

	fmt.Printf("Tasks by area:\n")
	fmt.Printf("Backend: %d\n", backend_cnt)
	fmt.Printf("English: %d\n", english_cnt)
	fmt.Printf("Algorithms: %d\n", algorithms_cnt)
	fmt.Printf("Guitar: %d\n", guitar_cnt)
	fmt.Printf("University: %d\n", university_cnt)
}
