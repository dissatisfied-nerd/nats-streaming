package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
)

var db *DBClient

func generateOutput(data interface{}) {
	datavalue := reflect.ValueOf(data)
	numOfFields := datavalue.NumField()
	dataType := datavalue.Type()

	for idx := 0; idx < numOfFields; idx++ {
		fieldValue := datavalue.Field(idx).Interface()
		fieldTag := dataType.Field(idx).Tag.Get("json")

		if len(fieldTag) > 0 {
			fmt.Printf("%s : %s\n", fieldTag, fieldValue)
		}
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[SERVER]: got request to root...")

	if r.URL.Query().Has("id") {
		id := r.URL.Query().Get("id")

		if id == "all" {
			data := db.GetAllOrders()

			outputData, _ := json.Marshal(data)
			io.WriteString(w, string(outputData[:]))
		} else {
			data, status := db.GetOrderById(id)

			if !status {
				io.WriteString(w, fmt.Sprintf("No order with id = %s", id))
			} else {
				outputData, _ := json.Marshal(data)
				io.WriteString(w, string(outputData[:]))
			}
		}
	}
}

func main() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_NAME")

	db = NewDbclient(dbUser, dbPassword, dbName)

	nsURL := os.Getenv("NATS_URL")
	nsCluster := os.Getenv("NATS_CLUSTER")
	nsClient := os.Getenv("NATS_SUBSCRIBER")
	nsChannel := os.Getenv("NATS_CHANNEL")

	ns := NewNSConnection(nsURL, nsCluster, nsClient)
	ns.Channel = nsChannel
	ns.Listen(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)

	err := http.ListenAndServe(":9090", mux)

	if err != nil {
		log.Fatalf("SERVER: %f", err)
	}

	db.GetAllOrders()
}
