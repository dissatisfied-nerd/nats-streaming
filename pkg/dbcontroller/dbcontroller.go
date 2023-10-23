package dbcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/dissatisfied-nerd/ns-service/pkg/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBClient struct {
	conn *sqlx.DB

	User     string
	Password string
	Name     string
}

// DBClient constructor
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

// generating SQL query
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

// Inseerting order
func (db *DBClient) InsertOrder(order model.Order) bool {
	if _, check := db.GetOrderById(order.Order_uid); check == true {
		return false
	}

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

	return true
}

// Selecting order by order_uid
func (db *DBClient) GetOrderById(id string) (model.Order, bool) {
	order := model.Order{}

	err := db.conn.Get(&order, "SELECT * FROM orders WHERE order_uid=$1", id)

	if order.Order_uid == "" {
		return order, false
	}

	if err != nil {
		log.Fatalf("DATABASE SELECT ORDER: %f", err)
	}

	err = db.conn.Get(&order.Payment, "SELECT * FROM payment WHERE order_id=$1", id)

	if err != nil {
		log.Fatalf("DATABASE SELECT PAYMENT: %f", err)
	}

	err = db.conn.Get(&order.Delivery, "SELECT * FROM delivery WHERE order_id=$1", id)

	if err != nil {
		log.Fatalf("DATABASE SELECT DELIVERY: %f", err)
	}

	items := model.Items{}

	err = db.conn.Get(&items, "SELECT * FROM items WHERE order_id=$1", id)

	if err != nil {
		log.Fatalf("DATABASE SELECT ITEMS: %f", err)
	}

	return order, true
}

// selecting all correct orders
func (db *DBClient) GetAllOrders() []model.Order {
	var result []model.Order
	orders := []model.Order{}

	err := db.conn.Select(&orders, "SELECT * FROM orders ORDER BY order_uid DESC")

	if err != nil {
		log.Fatalf("SELECT ALL: %f", err)
	}

	for _, order := range orders {
		err := db.conn.Get(&order.Payment, "SELECT * FROM payment WHERE order_id=$1", order.Order_uid)

		if err != nil {
			log.Fatalf("SELECT ALL: %f", err)
		}

		err = db.conn.Get(&order.Delivery, "SELECT * FROM delivery WHERE order_id=$1", order.Order_uid)

		if err != nil {
			log.Fatalf("SELECT ALL: %f", err)
		}

		item := model.Items{}

		err = db.conn.Get(&item, "SELECT * FROM items WHERE order_id=$1", order.Order_uid)

		if err != nil {
			log.Fatalf("SELECT ALL: %f", err)
		}

		order.Items = append(order.Items, item)

		result = append(result, order)
	}

	return result
}
