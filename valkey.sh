#!/usr/bin/bash

if [ -z "$1" ]; then
    IMAGE_TAG=""
else
    IMAGE_TAG=":$1"
fi

docker run \
    -d \
    -p 127.0.0.1:6379:6379 \
    --name "dev-valkey" \
    --rm \
    "valkey/valkey${IMAGE_TAG}"
