package httpapi

import (
	"strconv"
	"strings"

	"github.com/Zvoook/lifeforge/internal/task"

	"encoding/json"
	"net/http"
)

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	response := map[string]string{
		"status": "ok",
	}

	writeJSON(w, http.StatusOK, response)
}

func (s *Server) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetTasks(w, r)
	case http.MethodPost:
		s.handleCreateTask(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (s *Server) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := s.service.GetAllTasks()
	writeJSON(w, http.StatusOK, tasks)
}

func (s *Server) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var request CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeError(w, 400, "invalid request body")
		return
	}

	createdTask, err := s.service.CreateTask(request.Title, request.Area, request.Priority, request.EstimatedMinutes)
	if err != nil {
		writeError(w, 400, "invalid request body")
		return
	}

	err = task.SaveTasksToFile(s.service.GetAllTasks(), s.saveFileName)
	if err != nil {
		writeError(w, 500, "failed to save tasks")
		return
	}

	writeJSON(w, http.StatusOK, createdTask)
}

func (s *Server) handleTaskByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetTaskByID(w, r)
	// case http.MethodPost:
	// 	s.handleCreateTask(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (s *Server) handleGetTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	path := r.URL.Path
	path = strings.TrimPrefix(path, "/tasks/")
	id, err := strconv.Atoi(path)
	if err != nil {
		writeError(w, 400, "incorrect id")
		return
	}

	foundedTask, err := s.service.GetTaskById(uint32(id))
	if err == task.ErrTaskNotFound {
		writeError(w, 404, "task not found")
		return
	} else if err != nil {
		writeError(w, 500, "internal server error")
		return
	}
	writeJSON(w, 200, foundedTask)
}
