package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dissatisfied-nerd/nats-streaming/pkg/model"
	"github.com/nats-io/stan.go"
)

type NSConnection struct {
	conn stan.Conn

	URL     string
	Cluser  string
	Client  string
	Channel string
}

func NewNSConnection(nsUrl, nsCluster, nsClient string) *NSConnection {
	connection, err := stan.Connect(nsCluster, nsClient, stan.NatsURL(nsUrl))

	if err != nil {
		log.Fatal()
	}

	var NSConn NSConnection

	NSConn.conn = connection
	NSConn.URL = nsUrl
	NSConn.Cluser = nsCluster
	NSConn.Client = nsClient

	return &NSConn
}

func (ns *NSConnection) Listen(db *DBClient) {

	_, err := ns.conn.Subscribe(
		ns.Channel, func(msg *stan.Msg) {
			var order model.Order

			err := json.Unmarshal(msg.Data, &order)

			if err != nil {
				log.Fatalf("SUBSCRIBER: %f", err)
			}

			status := db.InsertOrder(order)

			if status {
				fmt.Printf("SUBSCRIBER: inserted order with order_uid = %s\n", order.Order_uid)
			}
		}, stan.DeliverAllAvailable())

	if err != nil {
		log.Fatalf("SUBSCRIBER: %s", err)
	}

}
