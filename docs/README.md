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
docker compose run -u root --rm parser python ./download_and_parse.py --year 2024 --output-path kdb_2024.json
mv ./parser/kdb_2024.json ./back/kdb_2024.json
docker compose run -u root --rm back go run .  update-courses-based-on-kdb --year 2024 --kdb-json-file-path kdb_2024.json
rm ./back/kdb_2024.json
```

start services
```sh
docker compose up db back front sponsorship proxy-docker
```

Access to http://localhost:4000 or http://localhost:4000/sponsorship

## Local Development - docker + host machine

Only proxy-host, db and db-migration are running in the docker container.

Version
- Go : 1.22.x
- Nodejs : nodejs 22.x.x
- Python : 3.12.x

Please install [direnv](https://github.com/direnv/direnv).
Remember to set [hook](https://direnv.net/docs/hook.html).

Example in Mac
```sh
brew install direnv
echo 'eval "$(direnv hook zsh)"' >> ~/.zshrc
source ~/.zshrc
```

Please Install [Bun](https://bun.sh/docs/installation).

Example in Mac
```sh
brew install oven-sh/bun/bun
```

Build Images
```sh
docker compose build db db-migration proxy-host
```

Terminal workspace
```sh
docker compose up -d db db-migration
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${TEST_DB_URL}'
```

Terminal parser
```sh
pip install -r requirements.txt
python download_and_parse.py --year 2024 --output-path kdb_2024.json
```

Terminal back
```sh
cd back

// setup environment variables
cp .env .env.local // please edit .env.local file
direnv allow .

// update courses based on kdb
go run .  update-courses-based-on-kdb --year 2024 --kdb-json-file-path ../parser/kdb_2024.json

// hot reload
go install github.com/air-verse/air@latest
air
```

Terminal front
```sh
cd front
bun install
bun run dev
```

Terminal sponsorship
```sh
cd sponsorship
bun install
bun run dev
```

Terminal workspace
```sh
docker compose up proxy-host
```

Access to http://localhost:4000 or http://localhost:4000/sponsorship
