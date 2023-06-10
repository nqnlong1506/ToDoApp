package databse

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() (*sql.DB, error) {
	// Database connection parameters
	connStr := "user=postgres password=nqnlong1506 dbname=todos sslmode=disable"

	// Establish the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the PostgreSQL database!")

	return db, nil
}

func InitlizeDatabase() {
	var err error
	DB, err = ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
