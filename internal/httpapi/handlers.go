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
		writeServiceError(w, err)
		return
	}

	createdTask, err := s.service.CreateTask(request.Title, request.Area, request.Priority, request.EstimatedMinutes)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	err = task.SaveTasksToFile(s.service.GetAllTasks(), s.saveFileName)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, 201, createdTask)
}

func (s *Server) handleTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetTask(w, r)
	case http.MethodPut:
		s.handleUpdateTask(w, r)
	case http.MethodDelete:
		s.handleDeleteTask(w, r)
	case http.MethodPatch:
		s.handleCompleteTask(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
}

func (s *Server) handleGetTask(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/tasks/")
	id, err := strconv.Atoi(path)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	foundedTask, err := s.service.GetTaskById(uint32(id))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, 200, foundedTask)
}

func (s *Server) handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/tasks/")
	id, err := strconv.Atoi(path)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	_, err = s.service.GetTaskById(uint32(id))
	if err != nil {
		writeServiceError(w, err)
		return
	}

	var request CreateTaskRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	updatedTask, err := s.service.UpdateTask(uint32(id), request.Title,
		request.Area, request.Status, request.Priority, request.EstimatedMinutes)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, 200, updatedTask)
}

func parsePath(path string, prefix string) (uint32, string, error) {
	path = strings.TrimPrefix(path, prefix)
	parts := strings.Split(path, "/")
	id, err := strconv.Atoi(parts[0])
	if len(parts) == 2 {
		action := parts[1]
		return uint32(id), action, err
	}
	return uint32(id), "", err
}

func (s *Server) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	id, _, err := parsePath(path, "/tasks/")
	if err != nil {
		writeServiceError(w, err)
		return
	}

	foundedTask, err := s.service.GetTaskById(uint32(id))
	if err == task.ErrTaskNotFound {
		writeServiceError(w, err)
		return
	} else if err != nil {
		writeServiceError(w, err)
		return
	}

	err = s.service.DeleteTask(uint32(id))
	if err != nil {
		writeServiceError(w, err)
		return
	}

	err = task.SaveTasksToFile(s.service.GetAllTasks(), s.saveFileName)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, 200, foundedTask)

}

func (s *Server) handleCompleteTask(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	id, action, err := parsePath(path, "/tasks/")
	if err != nil {
		writeServiceError(w, err)
		return
	} else if action != "complete" {
		writeError(w, 400, "incorrect action")
		return
	}

	foundedTask, err := s.service.GetTaskById(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	err = s.service.CompleteTask(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	foundedTask, err = s.service.GetTaskById(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	err = task.SaveTasksToFile(s.service.GetAllTasks(), s.saveFileName)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, 200, foundedTask)
}
