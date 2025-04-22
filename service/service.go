package service

import (
	"awesomeProject/types"
	"sync"
	"time"
)

// LayerService implements the business logic for handling requests
type LayerService struct {
	store types.RequestStore
	mu    sync.Mutex
}

// NewLayerService creates a new service layer with the provided store
func NewLayerService(store types.RequestStore) *LayerService {
	return &LayerService{
		store: store,
	}
}

// GetRequests retrieves all request records from the database
func (h *LayerService) GetRequests() ([]types.Request, error) {
	return h.store.GetRequests()
}

// CreateRequest creates a new request with default data and returns the generated ID
func (h *LayerService) CreateRequest() (int, error) {
	id, err := h.store.CreateRequest(types.Request{
		Name:   "Cola",
		Size:   500,
		Status: "Downloading",
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

// ProcessTask simulates a long-running background task
func (h *LayerService) ProcessTask(id int) {
	// Simulate processing delay (downloading a file, doing heavy computation)
	time.Sleep(3 * time.Minute)

	// Update the request as completed or failed
	h.mu.Lock()
	err := h.store.UpdateRequest(types.Request{
		Status:      "Completed",
		CompletedAt: time.Now(),
		ID:          id,
	})
	h.mu.Unlock()

	if err != nil {
		h.mu.Lock()
		_ = h.store.UpdateRequest(types.Request{
			Status:      "Failed",
			CompletedAt: time.Now(),
			ID:          id,
		})
		h.mu.Unlock()
		return
	}
}

// GetTaskStatus returns the current status of a specific task by ID
func (h *LayerService) GetTaskStatus(id int) (*types.Request, error) {
	var res *types.Request
	var err error

	h.mu.Lock()
	res, err = h.store.GetTaskStatus(id)
	h.mu.Unlock()

	return res, err
}
