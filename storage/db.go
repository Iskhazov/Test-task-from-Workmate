package storage

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

// NewMySQLStorage creates a new MySQL database connection using the provided configuration
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
