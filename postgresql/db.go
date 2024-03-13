package postgresql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() {
	host := os.Getenv("HOST")
	port := 5432
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"dbname=%s sslmode=disable", host, port, user, dbname)

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = database.Ping(); err != nil {
		panic(err)
	}

	db = database
}
