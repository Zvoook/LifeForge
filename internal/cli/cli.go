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
		case 9:
			c.clearAll()
		case 0:
			printInfo("Goodbye!")
			return
		default:
			printError(task.ErrUnknownAction)
		}

		c.waitForEnter()
	}
}
