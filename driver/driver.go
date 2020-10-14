package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var db *sql.DB

// ConnectDB : connects to the DB.
func ConnectDB() *sql.DB {
	// loads Environment variables in .env file
	gotenv.Load()

	// parse the postgres URL to fetch DB information.
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// Open the sql DB.
	db, err = sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}

	// Check DB connection has been established.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
