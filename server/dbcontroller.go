package main

import (
	"fmt"
	"log"

	"github.com/dissatisfied-nerd/nats-streaming/pkg/model"
	_ "github.com/lib/pq"

	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
)

type DBClient struct {
	conn *sqlx.DB

	User     string
	Password string
	Name     string
}

func NewDbclient(dbUser, dbPassword, dbName string) *DBClient {
	conn, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName))

	if err != nil {
		log.Fatalf("DATABASE: %f", err)
	}

	var db DBClient

	db.conn = conn

	db.User = dbUser
	db.Password = dbPassword
	db.Name = dbName

	db.conn.Ping()

	return &db
}

func generateQuery(tableName, data interface{}) string {
	query := fmt.Sprintf("INSERT INTO %s VALUES(", tableName)

	fields := structs.Fields(data)

	for _, field := range fields {
		tag := field.Tag("db")

		if tag != "" {
			query += fmt.Sprintf(":%s, ", tag)
		}
	}

	query = query[:len(query)-1] + ")"

	fmt.Println(query)

	return ""
}

func (db DBClient) InsertOrder(order model.Order) {
	generateQuery("delivery", order.Payment)
	generateQuery("delivery", order.Delivery)
	generateQuery("delivery", order)
}
