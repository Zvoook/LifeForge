package main

import (
	"encoding/json"
	"os"
)

func SaveTasksToFile(tasks []Task, filename string) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadTasksFromFile(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
	}

	var tasks []Task

	err = json.Unmarshal(data, &tasks)
	return tasks, nil
}
