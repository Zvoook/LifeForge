package main

import (
	"fmt"
	"strconv"
	"strings"
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
		return 0, ErrUnknownValue
	}

	return digit, nil
}

func (c *CLI) readID() (uint32, error) {
	id, err := c.readInt("Enter the ID: ")

	if err != nil {
		return 0, err
	}
	if id <= 0 {
		return 0, ErrInvalidId
	}

	return uint32(id), nil
}

func (c *CLI) readTitle() (string, error) {
	t, err := c.readLine("Enter the title: ")

	if err != nil {
		return "", err
	}
	if t == "" {
		return "", ErrInvalidTitle
	}

	return t, nil
}

func (c *CLI) readPriority() (int, error) {
	p, err := c.readInt("Enter the priority: ")

	if err != nil {
		return 0, err
	}
	if !validatePriority(p) {
		return 0, ErrInvalidPriority
	}

	return p, nil
}

func (c *CLI) readEstimatedMinutes() (int, error) {
	em, err := c.readInt("Enter estimated time in minutes: ")

	if err != nil {
		return 0, err
	}
	if !validateEstimatedMinutes(em) {
		return 0, ErrInvalidEstimatedMinutes
	}

	return em, nil
}

func (c *CLI) readArea() (Area, error) {
	printInfo("Areas: \n1. Backend \n2. English \n3. Guitar \n4. Algorithms \n5. University\n")
	a, err := c.readInt("Choose area: ")

	if err != nil {
		return Unknown, ErrInvalidArea
	}

	switch a {
	case 1:
		return Backend, nil
	case 2:
		return English, nil
	case 3:
		return Guitar, nil
	case 4:
		return Algorithms, nil
	case 5:
		return University, nil
	default:
		return Unknown, ErrInvalidArea
	}
}
