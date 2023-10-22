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

func generateQuery(tableName, data interface{}) (string, []interface{}) {
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

type queryData struct {
	query        string
	insertValues []interface{}
}

func (db *DBClient) InsertOrder(order model.Order) {
	orderQuery, orderValues := generateQuery("orders", order)
	paymentQuery, paymentValues := generateQuery("payment", order.Payment)
	deliveryQuery, deliveryValues := generateQuery("delivery", order.Delivery)

	var itemQueries []queryData

	for idx := range order.Items {
		var itemQuery queryData

		curQuery, curInsertValues := generateQuery("items", order.Items[idx])
		itemQuery.query = curQuery
		itemQuery.insertValues = curInsertValues

		itemQueries = append(itemQueries, itemQuery)
	}

	_, err := db.conn.Exec(orderQuery, orderValues...)

	if err != nil {
		log.Fatalf("DATABASE INSERT: %f", err)
	}

	_, err = db.conn.Exec(paymentQuery, paymentValues...)

	if err != nil {
		log.Fatalf("DATABASE INSERT: %f", err)
	}

	_, err = db.conn.Exec(deliveryQuery, deliveryValues...)

	if err != nil {
		log.Fatalf("DATABASE INSERT: %f", err)
	}

	for idx := range itemQueries {
		_, err := db.conn.Exec(itemQueries[idx].query, itemQueries[idx].insertValues...)

		if err != nil {
			log.Fatalf("DATABASE INSERT: %f", err)
		}
	}
}
