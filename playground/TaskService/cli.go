package main

import (
	"bufio"
	"os"
)

type CLI struct {
	Reader  *bufio.Reader
	Service TaskService
}

func NewCLI(s *TaskService) CLI {
	return CLI{Reader: bufio.NewReader(os.Stdin), Service: *s}
}

func (c *CLI) run() {
	for {
		clearScreen()
		printMenu()

		action, err := c.readInt("\nSelect action: ")

		if err != nil {
			clearScreen()
			printError(err)
			c.waitForEnter()
			continue
		}

		clearScreen()

		switch action {
		case 1:
			c.handleCreateTask()
		case 2:
			c.handleShowAllTasks()
		case 3:
			c.handleShowTasksByArea()
		case 4:
			c.handleFindTaskByID()
		case 5:
			c.handleCompleteTask()
		case 6:
			c.handleChangeTaskPriority()
		case 7:
			c.handleDeleteTask()
		case 8:
			c.showDashboard()
		case 0:
			printInfo("Goodbye!")
			return
		default:
			printInfo("Unknown action")
		}

		c.waitForEnter()
	}
}
