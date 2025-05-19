# Twin:te

## Local Development - only docker

build images
```sh
docker compose build db db-migration back front sponsorship proxy-docker parser codegen
```

environment variables on backend
```sh
cp back/.env back/.env.local // please edit it
```

db migration
```sh
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${TEST_DB_URL}'
```

update courses based on kdb
```sh
docker compose run -u root --rm parser python ./download_and_parse.py --year 2025 --output-path kdb_2025.json
mv ./parser/kdb_2025.json ./back/kdb_2025.json
docker compose run -u root --rm back go run .  update-courses-based-on-kdb --year 2025 --kdb-json-file-path kdb_2025.json
rm ./back/kdb_2025.json
```

start services
```sh
docker compose --profile docker up
```

Access to http://localhost:4000 or http://localhost:4000/sponsorship

## Local Development - docker + host machine

Only proxy-host, db and db-migration are running in the docker container.

Version
- Go : 1.23.x
- Nodejs : nodejs 22.x.x
- Python : 3.12.x

Please Install [Bun](https://bun.sh/docs/installation).

Example in Mac
```sh
brew install oven-sh/bun/bun
```

terminal root
```sh
# build images
docker compose build db db-migration proxy-host

# run containers
docker compose up -d db db-migration proxy-host

# run db migration
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${TEST_DB_URL}'
```

terminal parser
```sh
cd parser
pip install -r requirements.txt
python download_and_parse.py --year 2025 --output-path kdb_2025.json
```

terminal back
```sh
cd back

# setup environment variables
cp .env .env.local // please edit .env.local file
set -a; source .env.local; set +a;

# update courses based on kdb
go run .  update-courses-based-on-kdb --year 2025 --kdb-json-file-path ../parser/kdb_2025.json

# hot reload
go install github.com/air-verse/air@latest
air
```

terminal front
```sh
cd front
bun install
bun run dev
```

terminal sponsorship
```sh
cd sponsorship
bun install
bun run dev
```

Access to http://localhost:4000 or http://localhost:4000/sponsorship
