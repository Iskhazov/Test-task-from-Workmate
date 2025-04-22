package service

import (
	"awesomeProject/types"
	"awesomeProject/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Handler handles HTTP requests and delegates logic to the service layer
type Handler struct {
	service types.RequestService
}

// NewHandler initializes and returns a new Handler with the provided service
func NewHandler(service types.RequestService) *Handler {
	return &Handler{service: service}
}

// RegisterRoutes registers all endpoint routes to the provided router
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/requests", h.GetRequest).Methods(http.MethodGet)
	router.HandleFunc("/requests", h.NewRequest).Methods(http.MethodPost)
	router.HandleFunc("/requests/{taskID}", h.GetTaskStatus).Methods(http.MethodGet) // GET task status by task ID
}

// GetRequest handles GET method - returns all request records
func (h *Handler) GetRequest(w http.ResponseWriter, r *http.Request) {
	ps, err := h.service.GetRequests()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, ps)
}

// NewRequest handles POST method - creates a new request and triggers async processing
func (h *Handler) NewRequest(w http.ResponseWriter, r *http.Request) {
	id, err := h.service.CreateRequest()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, id)

	// Trigger asynchronous processing in a separate goroutine
	go h.service.ProcessTask(id)
}

// GetTaskStatus handles GET method - returns the status of a task by ID
func (h *Handler) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)["taskID"]
	id, err := strconv.Atoi(taskID) // Convert taskID string to integer
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.service.GetTaskStatus(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if task == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("task not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}
