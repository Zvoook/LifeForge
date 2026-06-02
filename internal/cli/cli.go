package cli

import (
	"bufio"
	"os"

	"github.com/Zvoook/lifeforge/internal/task"
)

type CLI struct {
	Reader       *bufio.Reader
	Service      *task.TaskService
	SaveFileName string
}

func NewCLI(service *task.TaskService, saveFileName string) CLI {
	return CLI{
		Reader:       bufio.NewReader(os.Stdin),
		Service:      service,
		SaveFileName: saveFileName,
	}
}

func (c *CLI) Run() {
	for {
		clearScreen()
		printMenu()

		action, err := c.readInt("\nSelect action: ")

		if err != nil {
			clearScreen()
			PrintError(err)
			c.waitForEnter()
			continue
		}

		clearScreen()

		switch action {
		case 1:
			c.handleShowAllTasks()
		case 2:
			c.handleShowTasksByArea()
		case 3:
			c.handleShowTasksByStatus()
		case 4:
			c.handleFindTaskByID()
		case 5:
			c.handleCreateTask()
		case 6:
			c.handleEditTask()
		case 7:
			c.handleCompleteTask()
		case 8:
			c.handleDeleteTask()
		case 9:
			c.showDashboard()
		case 10:
			c.clearAll()
		case 0:
			PrintInfo("Goodbye!")
			return
		default:
			PrintError(task.ErrUnknownAction)
		}

		c.waitForEnter()
	}
}
