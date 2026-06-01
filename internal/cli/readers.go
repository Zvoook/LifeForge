package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Zvoook/lifeforge/internal/task"
)

func (c *CLI) readLine(prompt string) (string, error) {
	fmt.Print(prompt)
	str, err := c.Reader.ReadString('\n')
	fmt.Print("\n")

	if err != nil {
		return "", err
	}

	str = strings.TrimSpace(str)

	return str, nil
}

func (c *CLI) readInt(prompt string) (int, error) {
	str, err := c.readLine(prompt)

	if err != nil {
		return 0, err
	}

	str = strings.TrimSpace(str)
	digit, err := strconv.Atoi(str)

	if err != nil {
		return 0, task.ErrUnknownValue
	}

	return digit, nil
}

func (c *CLI) readID() (uint32, error) {
	id, err := c.readInt("Enter the ID: ")

	if err != nil {
		return 0, err
	}
	if id <= 0 {
		return 0, task.ErrInvalidId
	}

	return uint32(id), nil
}

func (c *CLI) readTitle() (string, error) {
	t, err := c.readLine("Enter the title: ")

	if err != nil {
		return "", err
	}
	if t == "" {
		return "", task.ErrInvalidTitle
	}

	return t, nil
}

func (c *CLI) readPriority() (int, error) {
	p, err := c.readInt("Enter the priority: ")

	if err != nil {
		return 0, err
	}
	if !task.ValidatePriority(p) {
		return 0, task.ErrInvalidPriority
	}

	return p, nil
}

func (c *CLI) readEstimatedMinutes() (int, error) {
	em, err := c.readInt("Enter estimated time in minutes: ")

	if err != nil {
		return 0, err
	}
	if !task.ValidateEstimatedMinutes(em) {
		return 0, task.ErrInvalidEstimatedMinutes
	}

	return em, nil
}

func (c *CLI) readArea() (task.Area, error) {
	printInfo("Areas: \n1. Backend \n2. English \n3. Guitar \n4. Algorithms \n5. University\n")
	a, err := c.readInt("Choose area: ")

	if err != nil {
		return task.Unknown, task.ErrInvalidArea
	}

	switch a {
	case 1:
		return task.Backend, nil
	case 2:
		return task.English, nil
	case 3:
		return task.Guitar, nil
	case 4:
		return task.Algorithms, nil
	case 5:
		return task.University, nil
	default:
		return task.Unknown, task.ErrInvalidArea
	}
}
