package dbcontroller

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func NewDbclient(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	db.Ping()

	return db, nil
}
