package main

import (
	"fmt"
	"net/http"

	"github.com/Zvoook/lifeforge/internal/cli"
	"github.com/Zvoook/lifeforge/internal/httpapi"
	"github.com/Zvoook/lifeforge/internal/task"
)

const adress = ":8080"
const saveFileName = "save.json"

func main() {
	repository, err := task.NewRepositoryFromFile(saveFileName)
	if err != nil {
		fmt.Println("failed to initialise repository:", err)
		return
	}
	service := task.NewTaskService(&repository)

	server := httpapi.NewServer(&service, saveFileName)

	fmt.Println("LifeForge API is running on http://localhost:8080")

	err = http.ListenAndServe(adress, server.Routes())
	if err != nil {
		fmt.Println("failed to start server:", err)
		return
	}
}
