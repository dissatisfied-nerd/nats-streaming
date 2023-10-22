package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dissatisfied-nerd/nats-streaming/pkg/model"

	"github.com/nats-io/stan.go"
)

func NSConnect(nsUrl, nsCluster, nsClient string) stan.Conn {
	connection, err := stan.Connect(nsCluster, nsClient, stan.NatsURL(nsUrl))

	if err != nil {
		log.Fatal(err)
	}

	return connection
}

func ParseOrders(dataPath string) model.Order {
	file, err := os.ReadFile(dataPath)

	if err != nil {
		log.Fatal(err)
	}

	var orders model.Order

	err = json.Unmarshal(file, &orders)

	if err != nil {
		log.Fatal(err)
	}

	return orders
}

func main() {
	dataPath := os.Getenv("PUBLISHER_DATA_PATH")
	order := ParseOrders(dataPath)

	nsUrl := os.Getenv("NATS_URL")
	nsCluster := os.Getenv("NATS_CLUSTER")
	nsClient := os.Getenv("NATS_PUBLISHER")

	fmt.Printf("URL: %s, Cluster_ID: %s, Client_ID: %s \n", nsUrl, nsCluster, nsClient)
	fmt.Println("Connecting to nats-streaming-server...")

	connection := NSConnect(nsUrl, nsCluster, nsClient)

	natsChannel := os.Getenv("NATS_CHANNEL")

	var delay time.Duration = 10 * time.Second
	var tmp = 1

	for {
		order.Order_uid = fmt.Sprintf("%d", tmp)
		tmp++

		message, err := json.Marshal(order)

		if err != nil {
			log.Fatal(err)
		}

		err = connection.Publish(natsChannel, message)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sent message with order_uid = %s\n", order.Order_uid)

		time.Sleep(delay)
	}
}
