#!/bin/bash

main_dir="/home/axr/projects/nats-server"
publisher_dir="${main_dir}/publisher"
server_dir="${main_dir}/server"

export NATS_URL="localhost:8080"
export NATS_CLUSTER="cluster"
export NATS_SUBSCRIBER="client"
export NATS_PUBLISHER="publisher"
export NATS_CHANNEL="channel"

export PUBLISHER_DATA_PATH="/home/axr/projects/nats-server/publisher/data/model.json"

export POSTGRES_USER="postgres"
export POSTGRES_PASSWORD="postgres"
export POSTGRES_NAME="nats_streaming"

if [[ $1 == "publisher" ]]; then
    cd ${publisher_dir}
    go run .
fi

if [[ $1 == "server" ]]; then
    cd ${server_dir}
    go run .
fi

if [[ $1 == "all" ]]; then
    (docker run -p 8080:8080 -p 8223:8223 nats-streaming -p 8080 -m 8223 -cid cluster) &
    docker_pid=$!

    echo "docker pid = ${docker_pid}"

    sleep 1s

    cd ${publisher_dir}
    go run . &
    publisher_pid=$!

    echo "publisher pid = ${publisher_pid}"

    cd ${server_dir}
    go run . &
    server_pid=$!

    echo "server pid = ${server_pid}"

    while [[ ${signal} == "" ]]
    do
        read signal
    done

    sudo kill ${docker_pid}
    sudo kill ${publisher_pid}
    sudo kill ${server_pid}
fi


