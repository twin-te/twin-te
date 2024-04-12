## How to run

change directory

```sh
cd infra/dev
```

prepare environment variables

```sh
cp ../../back/.env.development ../../back/.env.local
```

please edit `../../back/.env.local` and configure OAuth2.0 (must, google is recommended) and stripe (optional)

run containers

```sh
docker compose up -d proxy
```

db migration

```sh
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${TEST_DB_URL}'
```

update courses based on KdB

```sh
docker compose run -u root --rm parser python ./download_and_parse.py --year 2024 --output-path kdb_2024.json
mv ../../parser/kdb_2024.json ../../back/kdb_2024.json
docker compose run -u root --rm back go run .  update-courses-based-on-kdb --year 2024 --kdb-json-file-path kdb_2024.json
rm ../../back/kdb_2024.json
```

access to http://localhost:8080

stop containers

```sh
docker compose stop
```
