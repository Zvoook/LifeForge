package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/Zvoook/lifeforge/internal/task"
)

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	response := errorResponse{
		Error: message,
	}

	writeJSON(w, statusCode, response)
}

func writeServiceError(w http.ResponseWriter, err error) {
	var serviceErrorStatusCodes = map[error]int{
		task.ErrTaskNotFound:            http.StatusNotFound,
		task.ErrTaskAlreadyCompleted:    http.StatusConflict,
		task.ErrInvalidTitle:            http.StatusBadRequest,
		task.ErrInvalidArea:             http.StatusBadRequest,
		task.ErrInvalidStatus:           http.StatusBadRequest,
		task.ErrInvalidPriority:         http.StatusBadRequest,
		task.ErrInvalidEstimatedMinutes: http.StatusBadRequest,
		task.ErrInvalidId:               http.StatusBadRequest,
	}

	statusCode, ok := serviceErrorStatusCodes[err]
	if !ok {
		writeError(w, http.StatusInternalServerError, "internal server error")
	}

	writeError(w, statusCode, err.Error())
}
