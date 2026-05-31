package main

import (
	"bufio"
	"os"
)

const Menu = "╔══════════════════════════════════════╗ \n║          LifeForge Task CLI          ║ \n╠══════════════════════════════════════╣ \n║ 1. Create task                       ║ \n║ 2. Show all tasks                    ║ \n║ 3. Show tasks by area                ║ \n║ 4. Find task by ID                   ║ \n║ 5. Complete task                     ║ \n║ 6. Change task priority              ║ \n║ 7. Delete task                       ║ \n║ 8. Show dashboard                    ║ \n║ 0. Exit                              ║ \n╚══════════════════════════════════════╝\n"

type CLI struct {
	Reader  *bufio.Reader
	Service TaskService
}

func NewCLI(s *TaskService) CLI {
	return CLI{Reader: bufio.NewReader(os.Stdin), Service: *s}
}

func (c *CLI) run() {
	for {
		c.printMenu()
		action, err := c.readInt("Select action:")

		if err != nil {
			c.printError(err)
			continue
		}

		switch action {
		case 1:
			err = c.handleCreateTask()
		case 2:
			c.handleShowAllTasks()
		case 3:
			err = c.handleShowTasksByArea()
		case 4:
			err = c.handleFindTaskByID()
		case 5:
			err = c.handleCompleteTask()
		case 6:
			err = c.handleChangeTaskPriority()
		case 7:
			err = c.handleDeleteTask()
		case 8:
			c.showDashboard()
		case 0:
			c.printInfo("Goodbye!")
			return
		default:
			c.printInfo("Unknown action")
		}

		if err != nil {
			c.printError(err)
			continue
		}
	}
}
