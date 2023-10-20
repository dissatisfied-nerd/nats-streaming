package dbcontroller

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func NewDbclient(*sqlx.Conn, error) {
	var DbURL = os.Getenv("POSTGRES_URL")

	db, err := sqlx.Connect("postgres", DbURL)

	if err != nil {
		log.Fatal(err)
	}

	db.Ping()
}
