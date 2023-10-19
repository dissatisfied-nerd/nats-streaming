#!/bin/bash

main_dir="/home/axr/projects/nats-server"
publisher_dir="${main_dir}/publisher"
server_dir="${main_dir}/server"

export NATS_URL="nats://user:password@localhost:4222"
export NATS_CLUSTER="test"
export NATS_CLIENT="test"
export NATS_CHANNEL="test"

export PUBLISHER_DATA_PATH="/home/axr/projects/nats-server/publisher/data/model.json"

export POSTGRES_URL="postgres://postgres:postgres@localhost:5432/nats_streaming"

if [[ $1 == "publisher" ]]; then
    cd ${publisher_dir}
fi

if [[ $1 == "server" ]]; then
    cd ${server_dir}
fi

go run main.go

