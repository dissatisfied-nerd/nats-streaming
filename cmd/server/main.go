package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dissatisfied-nerd/ns-service/pkg/cache"
	dbctl "github.com/dissatisfied-nerd/ns-service/pkg/dbcontroller"
	sub "github.com/dissatisfied-nerd/ns-service/pkg/subscriber"
)

var (
	mCache *cache.MemCache
	db     *dbctl.DBClient
)

// handler for http-server
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

// cache loader
func loadCache() {
	orders := db.GetAllOrders()

	for _, order := range orders {
		mCache.Add(order)
	}
}

func main() {
	//databse info
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_NAME")

	db = dbctl.NewDbclient(dbUser, dbPassword, dbName)

	//nats-streaming info
	nsURL := os.Getenv("NATS_URL")
	nsCluster := os.Getenv("NATS_CLUSTER")
	nsClient := os.Getenv("NATS_SUBSCRIBER")
	nsChannel := os.Getenv("NATS_CHANNEL")

	ns := sub.NewNSConnection(nsURL, nsCluster, nsClient)
	ns.Channel = nsChannel

	//loading cache from db
	mCache = cache.NewMemCache()
	loadCache()

	//listen asyncrhonously
	ns.Listen(db, mCache)

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)

	err := http.ListenAndServe(os.Getenv("SERVER_PORT"), mux)

	if err != nil {
		log.Fatalf("SERVER: %f", err)
	}
}
