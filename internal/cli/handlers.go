package cli

import (
	"fmt"

	"github.com/Zvoook/lifeforge/internal/task"
)

func (c *CLI) handleCreateTask() {
	title, err := c.readTitle()
	if err != nil {
		PrintError(err)
		return
	}

	area, err := c.readArea()
	if err != nil {
		PrintError(err)
		return
	}

	priority, err := c.readPriority()
	if err != nil {
		PrintError(err)
		return
	}

	em, err := c.readEstimatedMinutes()
	if err != nil {
		PrintError(err)
		return
	}

	createdTask, err := c.Service.CreateTask(title, area, priority, int(em))
	if err != nil {
		PrintError(err)
		return
	}

	err = task.SaveTasksToFile(c.Service.GetAllTasks(), c.SaveFileName)
	if err != nil {
		PrintError(err)
		return
	}

	PrintTaskActionResult("Task created\n", createdTask)
}

func (c *CLI) handleShowAllTasks() {
	tasks := c.Service.GetAllTasks()

	if len(tasks) == 0 {
		PrintInfo("No tasks found")
		return
	}

	printTasksTable(tasks)
}

func (c *CLI) handleShowTasksByArea() {
	tasks := c.Service.GetAllTasks()

	if len(tasks) == 0 {
		PrintInfo("Nothing to search")
		return
	}

	area, err := c.readArea()
	if err != nil {
		PrintError(err)
		return
	}

	tasks, err = c.Service.GetTasksByArea(area)
	if err != nil {
		PrintError(err)
		return
	}

	if len(tasks) == 0 {
		PrintInfo("No tasks found")
		return
	}

	printTasksTable(tasks)
}

func (c *CLI) handleShowTasksByStatus() {
	tasks := c.Service.GetAllTasks()

	if len(tasks) == 0 {
		PrintInfo("Nothing to search")
		return
	}

	status, err := c.readStatus()
	if err != nil {
		PrintError(err)
		return
	}

	tasks, err = c.Service.GetTasksByStatus(status)
	if err != nil {
		PrintError(err)
		return
	}

	if len(tasks) == 0 {
		PrintInfo("No tasks found")
		return
	}

	printTasksTable(tasks)
}

func (c *CLI) handleFindTaskByID() {
	id, err := c.selectIDForAction()
	if err != nil {
		PrintError(err)
	}

	foundTask, err := c.Service.GetTaskById(id)
	if err != nil {
		PrintError(err)
		return
	}

	printTasksTable([]task.Task{foundTask})
}

func (c *CLI) handleCompleteTask() {
	id, err := c.selectIDForAction()
	if err != nil {
		PrintError(err)
	}

	foundedTask, err := c.Service.GetTaskById(id)
	if err != nil {
		PrintError(err)
		return
	}

	err = c.Service.CompleteTask(id)
	if err != nil {
		PrintError(err)
		return
	}

	err = task.SaveTasksToFile(c.Service.GetAllTasks(), c.SaveFileName)
	if err != nil {
		PrintError(err)
		return
	}

	PrintTaskActionResult("Task completed\n", foundedTask)
}

func (c *CLI) handleEditTask() {
	id, err := c.selectIDForAction()
	if err != nil {
		PrintError(err)
	}

	beforeTask, err := c.Service.GetTaskById(uint32(id))
	if err != nil {
		PrintError(err)
		return
	}
	fmt.Printf("Task #%d: \n%v\n\n", id, beforeTask)

	par, err := c.readEditParameter()
	if err != nil {
		PrintError(err)
		return
	}

	switch par {
	case ParameterTitle:
		title, err := c.readTitle()
		if err != nil {
			PrintError(err)
			return
		}
		err = c.Service.EditTitle(id, title)
		if err != nil {
			PrintError(err)
			return
		}

	case ParameterArea:
		area, err := c.readArea()
		if err != nil {
			PrintError(err)
			return
		}
		err = c.Service.EditArea(id, area)
		if err != nil {
			PrintError(err)
			return
		}

	case ParameterStatus:
		status, err := c.readStatus()
		if err != nil {
			PrintError(err)
			return
		}
		err = c.Service.EditStatus(id, status)
		if err != nil {
			PrintError(err)
			return
		}

	case ParameterPriority:
		priority, err := c.readPriority()
		if err != nil {
			PrintError(err)
			return
		}
		err = c.Service.EditPriority(id, priority)
		if err != nil {
			PrintError(err)
			return
		}

	case ParameterEstimatedMinutes:
		minutes, err := c.readEstimatedMinutes()
		if err != nil {
			PrintError(err)
			return
		}
		err = c.Service.EditEstimatedMinutes(id, minutes)
		if err != nil {
			PrintError(err)
			return
		}
	}

	err = task.SaveTasksToFile(c.Service.GetAllTasks(), c.SaveFileName)
	if err != nil {
		PrintError(err)
		return
	}

	afterTask, err := c.Service.GetTaskById(uint32(id))
	if err != nil {
		PrintError(err)
		return
	}

	PrintTaskChangeResult("Task edited\n", beforeTask, afterTask)
}

func (c *CLI) handleDeleteTask() {
	id, err := c.selectIDForAction()
	if err != nil {
		PrintError(err)
	}

	deletedTask, err := c.Service.GetTaskById(id)

	err = c.Service.DeleteTask(id)
	if err != nil {
		PrintError(err)
		return
	}

	err = task.SaveTasksToFile(c.Service.GetAllTasks(), c.SaveFileName)
	if err != nil {
		PrintError(err)
		return
	}

	PrintSuccess("Task ")
	PrintTaskActionResult("Task deleted\n", deletedTask)
}

func (c *CLI) ShowDashboard() {
	tasks := c.Service.GetAllTasks()

	if len(tasks) == 0 {
		PrintInfo("Nothing to visualise")
		return
	}

	counterTaskByStatus := make(map[task.Status]int)
	counterTaskByArea := make(map[task.Area]int)
	totalTimeCount := 0
	highPriorityTaskCount := 0

	for _, oneTask := range tasks {
		totalTimeCount += oneTask.EstimatedMinutes

		if oneTask.Priority >= 8 {
			highPriorityTaskCount++
		}

		counterTaskByStatus[oneTask.Status]++
		counterTaskByArea[oneTask.Area]++
	}

	printDashboard(counterTaskByStatus, counterTaskByArea,
		len(tasks), totalTimeCount, highPriorityTaskCount)
}

func (c *CLI) ClearAll() {
	tasks := c.Service.GetAllTasks()

	if len(tasks) == 0 {
		PrintInfo("Nothing to clear")
		return
	}

	c.Service.ClearAll()
	err := task.SaveTasksToFile(c.Service.GetAllTasks(), c.SaveFileName)
	if err != nil {
		PrintError(err)
		return
	}
	PrintSuccess("All tasks cleared")
}

func (c *CLI) handleBuildDailyPlan() {
	timeLimit, err := c.readInt("Enter time limit: ")
	if err != nil {
		PrintError(err)
		return
	}

	tasks, err, totalTime := c.Service.BuildDailyPlan(timeLimit)
	if err != nil {
		PrintError(err)
	}

	if len(tasks) == 0 {
		PrintInfo("Nothing to plan")
		return
	}

	printForgeDailyPlan(tasks, timeLimit, totalTime)
}
