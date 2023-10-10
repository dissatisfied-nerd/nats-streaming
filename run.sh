#!/bin/bash

main_dir="/home/axr/projects/nats-server"
publisher_dir="${main_dir}/publisher"
subscriber_dir="${main_dir}/subscriber"

export NATS_URL="localhost:4040"
export NATS_CLUSTER="test-nats"
export NATS_CLIENT="test"
export NATS_CHANNEL="test"

export PUBLISHER_DATA_PATH="/home/axr/projects/nats-server/publisher/data/model.json"

cd ${publisher_dir}
go run main.go

