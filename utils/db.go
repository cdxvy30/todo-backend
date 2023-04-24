package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "postgres"
)

func ConnectToDB() (*sql.DB, error) {
	DBHOST := os.Getenv("DBHOST")
	if DBHOST == "" {
		DBHOST = "localhost"
	}
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DBHOST, USER, PASSWORD, DATABASE)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}

	return db, nil
}
