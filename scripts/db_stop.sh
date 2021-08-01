#!/bin/bash

if [ "$(docker ps -a -q -f name=suave-db)" ]; then
  docker rm -f suave-db
fi
