package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func NewDbclient(dbUser, dbPassword, dbName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName))

	if err != nil {
		log.Fatal(err)
	}

	db.Ping()

	return db, nil
}
