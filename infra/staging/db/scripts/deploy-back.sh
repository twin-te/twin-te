#!/bin/bash -eux

cd ~/twinte-infra-v4/staging/db

docker compose pull
docker compose build --no-cache
docker compose up -d
