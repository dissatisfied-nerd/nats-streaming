package dbcontroller

import (
	"log"
<<<<<<< HEAD
	"os"
=======
>>>>>>> server

	"github.com/jmoiron/sqlx"
)

<<<<<<< HEAD
func NewDbclient(*sqlx.Conn, error) {
	var DbURL = os.Getenv("POSTGRES_URL")

	db, err := sqlx.Connect("postgres", DbURL)
=======
func NewDbclient(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbURL)
>>>>>>> server

	if err != nil {
		log.Fatal(err)
	}

	db.Ping()
<<<<<<< HEAD
=======

	return db, nil
>>>>>>> server
}
