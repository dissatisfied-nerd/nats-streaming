CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS payment
(
    transaction   varchar(128),
    request_id    varchar(128),
    currency      varchar(8),
    provider      varchar(128),
    amount        int,
    payment_id    varchar(128),
    bank          varchar(128),
    delivery_cost int,
    goods_total   int,
    custom_fee    int
);

CREATE TABLE IF NOT EXISTS items
(
    chrt_id      int,
    track_number varchar(128),
    price        int,
    rid          varchar(128),
    name         varchar(128),
    sale         int,
    size         int,
    total_price  int,
    nm_id        int,
    brand        varchar(128),
    status       int
);

CREATE TABLE IF NOT EXISTS delivery
(
    name    varchar(128),
    phone   varchar(128),
    zip     varchar(128),
    city    varchar(128),
    address varchar(128),
    region  varchar(128),
    email   varchar(128)
);

CREATE TABLE IF NOT EXISTS orders
(
    locale             varchar(8),
    internal_signature varchar(128),
    customer_id        varchar(128),
    delivery_service   varchar(128),
    shardkey           varchar(128),
    sm_id              int,
    date_created       timestamp,
    off_shard          int,
 
    order_uid    varchar(128) PRIMARY KEY,
    track_number varchar(128),
    entry        varchar(128)
);

