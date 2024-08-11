#!/bin/bash -eux

cd ~/twinte-infra-v4/staging/app

docker compose pull
docker compose up -d
