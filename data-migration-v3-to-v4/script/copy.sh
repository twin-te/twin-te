#!/usr/bin/env bash

docker cp data-migration-v3-to-v4/data/processed twinte-db:/tmp/data
docker cp data-migration-v3-to-v4/script/migrate.sh twinte-db:/tmp