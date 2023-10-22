package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/dissatisfied-nerd/nats-streaming/pkg/model"
	_ "github.com/lib/pq"

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

func generateQuery(tableName, data interface{}) (string, interface{}) {
	var columns []string
	var variables []string
	var insertValues []interface{}

	dataValue := reflect.ValueOf(data)
	numOfFields := dataValue.NumField()
	dataType := dataValue.Type()

	for idx := 0; idx < numOfFields; idx++ {
		field := dataType.Field(idx)
		fieldValue := dataValue.Field(idx).Interface()

		if len(field.Tag.Get("db")) > 0 {
			var insertValue interface{}

			if len(field.Tag.Get("marsahl")) == 0 {
				insertValue = fieldValue
			} else {
				var err error

				insertValue, err = json.Marshal(fieldValue)

				if err != nil {
					log.Fatalf("DATABASE GENERATE QUERY: %f", err)
				}
			}

			columns = append(columns, field.Tag.Get("db"))
			variables = append(variables, fmt.Sprintf("$%d", len(columns)))
			insertValues = append(insertValues, insertValue)
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName, strings.Join(columns, ", "), strings.Join(variables, ", "))

	return query, insertValues
}

func (db *DBClient) InsertOrder(order model.Order) {
	query, insertValues := generateQuery("orders", order)

	_, err := db.conn.Exec(query, insertValues)

	if err != nil {
		log.Fatalf("DATABASE INSERT: %f", err)
	}
}
