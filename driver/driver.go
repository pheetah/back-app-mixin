package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANT_SQL_URL"))

	if err != nil {
		log.Fatal("Error occured while opening the database", err)
	}

	db, err = sql.Open("postgres", pgUrl)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return db
}
