#!/bin/bash

# NOTE: docker compose up db などで db が立ち上がっている状態で実行すること

# COMPOSE_FILE 環境変数から参照（未設定の場合はデフォルトのパスを使用）
COMPOSE_FILE=${COMPOSE_FILE:-../docker-compose.yml}

rm -rf data/processed/*

pip install -r requirements.txt
pip install -r ../parser/requirements.txt

mkdir -p data/parsed
years=(2019 2020 2021 2022 2023 2024)
for year in "${years[@]}"; do
  python3 ../parser/download_and_parse.py --year "$year" --output-path data/parsed/"$year".json
done

mkdir -p data/processed
python3 main.py

# db-migration 用のコンテナ起動（-f で参照するファイルは環境変数 COMPOSE_FILE を利用）
docker compose -f "$COMPOSE_FILE" run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'

# テーブル名と対応するCSVファイルの組み合わせ
csv_groups=(
  "courses:courses_found.csv courses_not_found.csv"  
  "course_methods:course_methods_found.csv course_methods_not_found.csv"
  "course_recommended_grades:course_recommended_grades_found.csv course_recommended_grades_not_found.csv"
  "course_schedules:course_schedules_found.csv course_schedules_not_found.csv"
  "users:users.csv"
  "user_authentications:user_authentications.csv"
  "sessions:sessions.csv"
  "payment_users:payment_users.csv"
  "tags:tags.csv"
  "registered_courses:registered_courses.csv"
  "registered_course_tag_ids:registered_course_tag_ids.csv"
)

# コンテナ内の一時ディレクトリ作成（サービス名を db に変更）
docker compose -f "$COMPOSE_FILE" exec db mkdir -p /tmp/v3_dump

# 全てのCSVファイルをコンテナにコピー
for group in "${csv_groups[@]}"; do
  IFS=":" read -r table csvs <<< "$group"
  for csv in $csvs; do
    # docker compose cp を利用（サービス名は db）
    docker compose cp data/processed/"$csv" db:/tmp/v3_dump/"$csv"
  done
done

# PostgreSQL の COPY コマンドで "null" を文字列と解釈しないように変換
# Ref: https://www.postgresql.org/docs/current/sql-copy.html
docker compose -f "$COMPOSE_FILE" exec db sh -c "sed -i 's/\"null\"/null/g' /tmp/v3_dump/*"

POSTGRES_URL=${POSTGRES_URL:-"postgres://postgres:password@db:5432/twinte_db?sslmode=disable"}

# 各テーブルにデータをインポート
for group in "${csv_groups[@]}"; do
  IFS=":" read -r table csvs <<< "$group"
  for csv in $csvs; do
    echo "Importing $csv to $table"
    docker compose exec -T db sh -c "psql -d $POSTGRES_URL -c \"COPY $table FROM '/tmp/v3_dump/$csv' WITH (FORMAT csv, HEADER true, NULL 'NULL')\""
  done
done
