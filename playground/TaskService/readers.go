package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (c *CLI) readLine(prompt string) (string, error) {
	fmt.Println(prompt)
	str, err := c.Reader.ReadString('\n')

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
		return 0, err
	}

	return digit, nil
}

func (c *CLI) readID() (uint32, error) {
	id, err := c.readInt("Enter the ID:")

	if err != nil {
		return 0, err
	}
	if id <= 0 {
		return 0, ErrInvalidId
	}

	return uint32(id), nil
}

func (c *CLI) readTitle() (string, error) {
	t, err := c.readLine("Enter the title:")

	if err != nil {
		return "", err
	}
	if t == "" {
		return "", ErrInvalidTitle
	}

	return t, nil
}

func (c *CLI) readPriority() (uint8, error) {
	p, err := c.readInt("Enter the priority:")

	if err != nil {
		return 0, err
	}
	if !validatePriority(uint8(p)) {
		return 0, ErrInvalidPriority
	}

	return uint8(p), nil
}

func (c *CLI) readEstimatedMinutes() (uint32, error) {
	em, err := c.readInt("Enter estimated time in minutes:")

	if err != nil {
		return 0, err
	}
	if !validateEstimatedMinutes(uint32(em)) {
		return 0, ErrInvalidEstimatedMinutes
	}

	return uint32(em), nil
}

func (c *CLI) readArea() (Area, error) {
	a, err := c.readInt("Choose area: \n1. Backend \n2. English \n3. Guitar \n4. Algorithms \n5. University")

	if err != nil {
		return Unknown, err
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
