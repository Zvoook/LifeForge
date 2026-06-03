package cli

import (
	"fmt"

	"github.com/Zvoook/lifeforge/internal/task"
)

func (c *CLI) selectIDForAction() (uint32, error) {
	tasks := c.Service.GetAllTasks()

	if len(tasks) == 0 {
		return 0, task.ErrNothingToChange
	} else {
		fmt.Print("Available tasks: \n")
		printTasksTable(tasks)
		fmt.Print("\n")
	}

	id, err := c.readID()
	if err != nil {
		return 0, err
	}

	return id, nil
}
