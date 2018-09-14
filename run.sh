#!/bin/bash

rm -r /tmp/nats

pushd nats-streaming
docker build -t nats-streaming-exp .
popd

pushd pub
docker build -t nats-pub .
popd

pushd sub
docker build -t nats-sub .
popd

blockade destroy
blockade up

sleep 30

while true; do
    echo "partitioning"
    blockade partition nats-streaming
    sleep $((15+($RANDOM%25)))
    blockade kill nats-streaming

    sleep 15

    echo "starting nats-streaming and subscriber"
    blockade join
    blockade start nats-streaming
    blockade start pub
    blockade start sub_1
    blockade start sub_2
    blockade start sub_3
    blockade start sub_4
    blockade start sub_5
    blockade start sub_6
    blockade start sub_7
    blockade start sub_8
    blockade start sub_9
    blockade start sub_10

    sleep 5

    echo "logs:"
    blockade logs sub_1
    blockade logs sub_2
    blockade logs sub_3
    blockade logs sub_4
    blockade logs sub_5
    blockade logs sub_6
    blockade logs sub_7
    blockade logs sub_8
    blockade logs sub_9
    blockade logs sub_10
done