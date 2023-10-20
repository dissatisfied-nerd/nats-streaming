#!/bin/bash

main_dir="/home/axr/projects/nats-server"
publisher_dir="${main_dir}/publisher"
server_dir="${main_dir}/server"

export NATS_URL="localhost:8080"
export NATS_CLUSTER="test"
export NATS_CLIENT="test"
export NATS_CHANNEL="test"

export PUBLISHER_DATA_PATH="/home/axr/projects/nats-server/publisher/data/model.json"

export POSTGRES_USER="postgres"
export POSTGRES_PASSWORD="postgres"
export POSTGRES_NAME="nats_streaming"

if [[ $1 == "publisher" ]]; then
    cd ${publisher_dir}
fi

if [[ $1 == "server" ]]; then
    cd ${server_dir}
fi

go run .

