package dbmodel

import "time"

type Payment struct {
	id int

	transaction   string
	request_id    string
	currency      string
	provider      string
	amount        int
	payment_id    string
	bank          string
	delivery_cost int
	goods_total   int
	custom_fee    int
}

type Items struct {
	chrt_id      string
	track_number string
	price        int
	rid          string
	name         string
	sale         int
	size         int
	total_price  int
	nm_id        int
	brand        string
	status       int
}

type Delivery struct {
	id int

	name    string
	phone   string
	zip     string
	city    string
	address string
	region  string
	email   string
}

type Specs struct {
	id int

	locale             string
	internal_signature string
	customer_id        string
	delivery_service   string
	shardkey           string
	sm_id              int
	date_created       time.Time
	off_shard          int
}

type Orders struct {
	Id int

	order_uid    string
	track_number string
	entry        string

	Delivery Delivery
	Payment  Payment
	Items    []Items
	Specs    Specs
}
