#!/usr/bin/bash

if [ -z "$1" ]; then
    IMAGE_TAG=""
else
    IMAGE_TAG=":$1"
fi

docker run \
    -d \
    -e MYSQL_ROOT_PASSWORD="password" \
    -e MYSQL_ROOT_HOST='%' \
    --name "dev-mysql" \
    -p 127.0.0.1:3306:3306 \
    --rm \
    "mysql${IMAGE_TAG}"
