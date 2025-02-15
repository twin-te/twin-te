sudo docker compose -f docker-compose.yml -f docker-compose.override.yml run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'

# コンテナに dump をコピー
find -type f -name "*.dump" | xargs -I{} bash -c 'sudo docker cp {} twinte-db:/root'

# データのリストア
# --data-only: データのみをリストア
# --disable-triggers: トリガーを無効化（エラーがでるため）
find -type f -name "*.dump" | xargs -I{} sudo docker exec -i twinte-db sh -c 'PGPASSWORD="password" pg_restore --data-only --disable-triggers -d postgres://postgres:password@db:5432/twinte_db?sslmode=disable /root/$(basename {})'

# 既存データのマイグレーション
sudo docker exec twinte-db psql -d "postgres://postgres:password@db:5432/twinte_db?sslmode=disable" -f /root/onetime/20240902_flatten_annual.sql
