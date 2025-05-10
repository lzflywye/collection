#!/usr/bin/bash

if [ -z "$1" ]; then
    IMAGE_TAG=""
else
    IMAGE_TAG=":$1"
fi

docker run \
    -d \
    -e POSTGRES_PASSWORD=password \
    --name "dev-postgres" \
    -p 127.0.0.1:5432:5432 \
    --rm \
    "postgres${IMAGE_TAG}"
