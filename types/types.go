package types

import "time"

// RequestStore defines the methods required for interacting with the database layer.
type RequestStore interface {
	GetRequests() ([]Request, error)        // Retrieve all requests
	CreateRequest(Request) (int, error)     // Create a new request and return its ID
	UpdateRequest(Request) error            // Update an existing request (status and time)
	GetTaskStatus(id int) (*Request, error) // Get the status of a specific task by ID
}

// RequestService defines the methods required for service layer.
type RequestService interface {
	GetRequests() ([]Request, error)
	CreateRequest() (int, error)
	ProcessTask(id int) // Asynchronously process the task
	GetTaskStatus(id int) (*Request, error)
}

// Request represents a request record in the system (download, task, etc.)
type Request struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Size        int       `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
	Status      string    `json:"status"` // Status: Downloading, Completed or Failed.
}

// NewRequest is for creating a new request via API input.
type NewRequest struct {
	Name   string `json:"name" validate:"required"`
	Size   int    `json:"size"`
	Status string `json:"quantity" validate:"required"`
}
