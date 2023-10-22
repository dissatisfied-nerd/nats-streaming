CREATE TABLE IF NOT EXISTS orders
(
    order_uid    varchar(128) PRIMARY KEY,
    track_number varchar(128),
    entry        varchar(128),

    locale             varchar(8),
    internal_signature varchar(128),
    customer_id        varchar(128),
    delivery_service   varchar(128),
    shardkey           varchar(128),
    sm_id              int,
    date_created       timestamp,
    off_shard          int
);

CREATE TABLE IF NOT EXISTS payment
(
    order_id varchar(128) PRIMARY KEY, 
        
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
    order_id varchar(128) PRIMARY KEY,

    track_number varchar(128),
    chrt_id      int,
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
    order_id varchar(128) PRIMARY KEY,

    name    varchar(128),
    phone   varchar(128),
    zip     varchar(128),
    city    varchar(128),
    address varchar(128),
    region  varchar(128),
    email   varchar(128)
);


