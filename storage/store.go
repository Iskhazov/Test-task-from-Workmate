package storage

import (
	"awesomeProject/types"
	"database/sql"
	"log"
)

// Store is a concrete implementation of the RequestStore interface
type Store struct {
	db *sql.DB
}

// NewStore returns a new instance of Store using the provided database connection
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetRequests retrieves all request records from the database
func (s *Store) GetRequests() ([]types.Request, error) {
	rows, err := s.db.Query("SELECT * FROM requests")
	if err != nil {
		return nil, err
	}

	requests := make([]types.Request, 0)
	for rows.Next() {
		p, err := scanRowsIntoRequests(rows)
		if err != nil {
			return nil, err
		}
		requests = append(requests, *p)
	}

	return requests, nil
}

// CreateRequest inserts a new request record into the database
func (s *Store) CreateRequest(request types.Request) (int, error) {
	_, err := s.db.Exec("INSERT INTO requests(name, size, status) VALUES(?,?,?)",
		request.Name, request.Size, request.Status)
	if err != nil {
		return 0, err
	}

	// Get the Id
	var id int
	err = s.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

// UpdateRequest updates the status and set completion time of an existing request
func (s *Store) UpdateRequest(request types.Request) error {
	_, err := s.db.Exec("UPDATE requests SET status = ?, completed_at = ? WHERE id = ?",
		request.Status, request.CompletedAt, request.ID)
	return err
}

// GetTaskStatus fetches the request by its Id and returns its current state
func (s *Store) GetTaskStatus(id int) (*types.Request, error) {
	row, err := s.db.Query("SELECT * FROM requests WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	var p *types.Request
	for row.Next() {
		p, err = scanRowsIntoRequests(row)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

// scanRowsIntoRequests maps the database columns from a single row into a Request struct
func scanRowsIntoRequests(rows *sql.Rows) (*types.Request, error) {
	request := new(types.Request)
	err := rows.Scan(
		&request.ID,
		&request.Name,
		&request.Size,
		&request.Status,
		&request.CreatedAt,
		&request.CompletedAt,
	)
	if err != nil {
		return nil, err
	}
	return request, nil
}
