# Local Environment

## How to run

change directory

```sh
cd infra/local
```

prepare environment variables

```sh
cp ../../back/.env ../../back/.env.local
```

please edit `../../back/.env.local` and configure OAuth2.0 (required, google is recommended) and stripe (optional)

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

access to http://localhost:8080 or http://localhost:8080/sponsorship

stop containers

```sh
docker compose stop
```

## Useful commands

start bash in back container

```sh
docker compose exec -it back bash
```

start psql in db container
```sh
docker compose exec -it db psql -U postgres -d twinte_db
```

remove all containers

```sh
docker rm -f $(docker ps -aq)
```

remove unused data

```sh
docker system prune --all --force --volumes
```
