package main

import (
	"awesomeProject/config"
	"awesomeProject/storage"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func main() {
	// Create MySQL connection config using environment variables
	db, err := storage.NewMySQLStorage(mysqlCfg.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp", // Connection protocol
		AllowNativePasswords: true,  // Allow native authentication
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Initialize migration driver with an already opened DB connection
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a migration instance, specify path to migration files and database name
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Read the last command-line argument (expecting "up" or "down")
	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up" {
		// Create table
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		// Drop table
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
