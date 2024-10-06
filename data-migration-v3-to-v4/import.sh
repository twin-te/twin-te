#!/bin/bash

# NOTE: docker compose up db などで db が立ち上がっている状態で実行すること

rm -rf data/processed/*

pip install -r requirements.txt

python3 main.py

docker compose -f ../docker-compose.yml run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'

find -type f -path "./data/processed/*.csv" | xargs -I{} bash -c 'docker cp {} twinte-db:/root'

find -type f -path "./data/processed/*.csv" | xargs -I{} docker exec -i twinte-db sh -c '
  FILENAME=$(basename {})
  TABLE_NAME="${FILENAME%.*}"
  # TODO: 本番移行時は良い感じに取得する
  POSTGRES_URL="postgres://postgres:password@db:5432/twinte_db?sslmode=disable"
  psql -d $POSTGRES_URL -c "\COPY $TABLE_NAME FROM /root/$FILENAME WITH CSV HEADER"
'
