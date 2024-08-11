#!/bin/bash -eux

cd $(dirname $0) || exit
cd ../

docker compose pull
docker compose up -d
