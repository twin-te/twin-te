#!/bin/bash

# NOTE: docker compose up db などで db が立ち上がっている状態で実行すること

rm -rf data/processed/*

pip install -r requirements.txt
pip install -r ../parser/requirements.txt

mkdir -p data/parsed
years=(2019 2020 2021 2022 2023 2024)
for year in ${years[@]}; do
  python3 ../parser/download_and_parse.py --year $year --output-path data/parsed/$year.json
done

mkdir -p data/processed
python3 main.py

docker compose -f ../docker-compose.yml run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'

# テーブル名と対応するCSVファイルの組み合わせ
csv_groups=(
  "course_methods:course_methods_found.csv course_methods_not_found.csv"
  "course_recommended_grades:course_recommended_grades_found.csv course_recommended_grades_not_found.csv"
  "course_schedules:course_schedules_found.csv course_schedules_not_found.csv"
  "courses:courses_found.csv courses_not_found.csv"
  "payment_users:payment_users.csv"
  "registered_courses:registered_courses.csv"
  "registered_course_tag_ids:registered_course_tag_ids.csv"
  "sessions:sessions.csv"
  "tags:tags.csv"
  "user_authentications:user_authentications.csv"
  "users:users.csv"
)

docker exec twinte-db sh -c "mkdir -p /tmp/v3_dump"

# 全てのCSVファイルをコンテナにコピー
for group in "${csv_groups[@]}"; do
  IFS=":" read -r table csvs <<< "$group"
  for csv in $csvs; do
    docker cp data/processed/$csv twinte-db:/tmp/v3_dump/$csv
  done
done

# TODO: 本番移行時は良い感じに取得する
POSTGRES_URL="postgres://postgres:password@db:5432/twinte_db?sslmode=disable"

# 各テーブルにデータをインポート
for group in "${csv_groups[@]}"; do
  IFS=":" read -r table csvs <<< "$group"
  for csv in $csvs; do
    docker exec -i twinte-db sh -c "psql -d $POSTGRES_URL -c \"COPY $table FROM '/tmp/v3_dump/$csv' WITH CSV HEADER\""
  done
done
