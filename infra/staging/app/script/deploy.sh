#!/bin/bash -eux

cd $(dirname $0) || exit
cd ../

git switch master
git pull

docker compose pull
docker compose up -d
