package main

import (
	"fmt"
	"log"
	"reflect"

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

func generateQuery(tableName, data interface{}) string {
	query := fmt.Sprintf("INSERT INTO %s (", tableName)

	value := reflect.ValueOf(data)
	numOfFields := value.NumField()
	strcutType := value.Type()

	for i := 0; i < numOfFields; i++ {
		field := strcutType.Field(i)
		tag := field.Tag.Get("db")

		if len(tag) > 0 {
			query += fmt.Sprintf("%s,", tag)
		}
	}

	query = query[:len(query)-1] + ") VALUES ("

	for i := 0; i < numOfFields; i++ {
		field := strcutType.Field(i)
		tag := field.Tag.Get("db")

		if len(tag) > 0 {
			query += fmt.Sprintf(":%s,", tag)
		}
	}

	query = query[:len(query)-1] + ")"

	return query
}

func (db *DBClient) InsertOrder(order model.Order) {
	orderQuery := generateQuery("order", order)
	paymentQuery := generateQuery("payment", order.Payment)
	deliveryQuery := generateQuery("delivery", order.Delivery)

	var itemsQuery []string

	for idx := range order.Items {
		itemsQuery = append(itemsQuery, generateQuery("items", order.Items[idx]))
	}

	fmt.Println(orderQuery)

	_, err := db.conn.NamedExec(orderQuery, order)

	if err != nil {
		fmt.Println(err)
		log.Fatalf("DATABASE INSERT: %f", err)
	}

	_, err = db.conn.NamedExec(paymentQuery, order.Payment)

	if err != nil {
		log.Fatalf("DATABASE INSERT: %f", err)
	}

	_, err = db.conn.NamedExec(deliveryQuery, order.Delivery)

	if err != nil {
		log.Fatalf("DATABASE INSERT: %f", err)
	}

	for idx := range itemsQuery {
		_, err = db.conn.NamedExec(itemsQuery[idx], order.Items[idx])

		if err != nil {
			log.Fatalf("DATABASE INSERT: %f", err)
		}
	}
}
