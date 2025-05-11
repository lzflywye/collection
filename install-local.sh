#!/usr/bin/bash

if [ -z "$1" ]; then
  echo "$0 <binary>"
  exit 1
fi

install -t /usr/local/bin/ -o root -g root -m 755 "$1"
