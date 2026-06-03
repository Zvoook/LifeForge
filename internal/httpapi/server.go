package httpapi

import (
	"net/http"

	"github.com/Zvoook/lifeforge/internal/task"
)

type Server struct {
	service      *task.TaskService
	saveFileName string
}

func NewServer(service *task.TaskService, saveFileName string) *Server {
	return &Server{
		service:      service,
		saveFileName: saveFileName,
	}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/tasks", s.handleTasks)
	mux.HandleFunc("/tasks/", s.handleTaskByID)

	return mux
}
