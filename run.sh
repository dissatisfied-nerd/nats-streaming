#!/bin/bash

main_dir="/home/axr/projects/nats-server"
publisher_dir="${main_dir}/publisher"
subscriber_dir="${main_dir}/subscriber"

export NATSURL="localhost:4040"
export NATSCLUSTER="test-nats"
export NATSCLIENT="test"
export NATSCHANNEL="test"

go run publisher/main.go

