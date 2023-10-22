package model

import "time"

type Payment struct {
	Transaction   string `db:"transaction" json:"transaction"`
	Request_id    string `db:"request_id" json:"request_id"`
	Currency      string `db:"currency" json:"currency"`
	Provider      string `db:"provider" json:"provider"`
	Amount        int    `db:"amount" json:"amount"`
	Payment_id    string `db:"payment_id" json:"payment_id"`
	Bank          string `db:"bank" json:"bank"`
	Delivery_cost int    `db:"delivery_cost" json:"delivery_cost"`
	Goods_total   int    `db:"goods_total" json:"goods_total"`
	Custom_fee    int    `db:"custom_fee" json:"custom_fee"`
}

type Items struct {
	Chrt_id      int    `db:"chrt_id" json:"chrt_id"`
	Track_number string `db:"track_number" json:"track_number"`
	Price        int    `db:"price" json:"prive"`
	Rid          string `db:"rid" json:"rid"`
	Name         string `db:"name" json:"name"`
	Sale         int    `db:"sale" json:"sale"`
	Size         string `db:"size" json:"size"`
	Total_price  int    `db:"total_price" json:"total_price"`
	Nm_id        int    `db:"nm_id" json:"nm_id"`
	Brand        string `db:"brand" json:"brand"`
	Status       int    `db:"status" json:"status"`
}

type Delivery struct {
	Name    string `db:"name" json:"name"`
	Phone   string `db:"phone" json:"phone"`
	Zip     string `db:"zip" json:"zip"`
	City    string `db:"city" json:"city"`
	Address string `db:"address" json:"address"`
	Region  string `db:"region" json:"region"`
	Email   string `db:"email" json:"email"`
}

type Order struct {
	Locale             string    `db:"locale" json:"locale"`
	Internal_signature string    `db:"internal_signature" json:"internal_signature"`
	Customer_id        string    `db:"customer_id" json:"customer_id"`
	Delivery_service   string    `db:"delivery_service" json:"delivery_service"`
	Shardkey           string    `db:"shardkey" json:"shardkey"`
	Sm_id              int       `db:"sm_id" json:"sm_id"`
	Date_created       time.Time `db:"date_created" json:"date_created"`
	Off_shard          int       `db:"off_shard" json:"off_shard"`

	Order_uid    string `db:"order_uid" json:"order_uid"`
	Track_number string `db:"track_number" json:"track_number"`
	Entry        string `db:"entry" json:"entry"`

	Delivery Delivery `json:"delivery" marshal:"1"`
	Payment  Payment  `json:"payment" marshal:"2"`
	Items    []Items  `json:"items" marshal:"3"`
}
