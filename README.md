# Twin:te

## Local Development - only docker

Please see [infra/local](./infra/local/README.md)

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

Terminal workspace
```sh
docker compose up -d db db-migration
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${TEST_DB_URL}'
```

Terminal parser
```sh
pip install -r requirements.txt
python download_and_parse.py --year 2024 --output-path kdb.json
```

Terminal back
```sh
cd back

// setup environment variables
cp .env .env.local // please edit .env.local file
direnv allow .

// download course data from kdb
go run .  update-courses-based-on-kdb --year 2024 --kdb-json-file-path ../parser/kdb.json

// hot reload
go install github.com/air-verse/air@latest
asdf reshim golang
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

Access to http://localhost or http://localhost/sponsorship
