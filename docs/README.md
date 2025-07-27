# Twin:te

## 開発環境の構築方法（Docker）

Twin:te をローカルで動かすための最低限の手順を以下に示します。

### 環境変数の設定

`.env.local` を作成します。

```sh
cp ./back/.env ./back/.env.local
```

Twin:te をローカルで動かすためには最低限 Google OAuth2.0 の設定が必要です。
`.env.local` の `OAUTH_GOOGLE_CLIENT_ID` と `OAUTH_GOOGLE_CLIENT_SECRET` にそれぞれ取得した情報を設定してください。
Twin:te 関係者は共有されている環境変数を参照できます。
外部コントリビュータの方は[Google OAuth2.0 の設定](./setup-google-oauth/README.md)で入手した情報を用いてください。

### 立ち上げ

最初に DB のマイグレーションをします。

```sh
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${TEST_DB_URL}'
```

次に [KdB](https://kdb.tsukuba.ac.jp/) から最新の講義情報を取得します。

```sh
docker compose run -u root --rm parser uv run ./download_and_parse.py --year 2025 --output-path kdb_2025.json
mv ./parser/kdb_2025.json ./back/kdb_2025.json
docker compose run -u root --rm back go run . update-courses-based-on-kdb --year 2025 --kdb-json-file-path kdb_2025.json
rm ./back/kdb_2025.json
```

アプリケーションを立ち上げます。

```sh
docker compose --profile docker up
```

`http://localhost:4000` で Twin:te が使用できます。

## 寄付ページの開発

寄付ページをローカルで動かすための手順を以下に示します。

### 環境変数の設定

`/back/.env.local` の `STRIPE_KEY` に Stripe の API キーを設定してください。
Twin:te 関係者は共有されている環境変数を参照できます。外部コントリビュータの方は[Stripe の設定](./setup-stripe/README.md)で入手した情報を用いてください。

### 立ち上げ

アプリケーションを立ち上げます。

```sh
docker compose --profile docker up
```

`http://localhost:4000/sponsorship` で寄付ページが使用できます。

## 開発環境の構築方法（Docker + ホストマシン）

proxy-host, db, db-migrationのみをDockerコンテナで実行し、他のサービスはホストマシン上で実行する方法です。

バージョン
- Go : 1.23.x
- Python : 3.12.x
- [uv](https://docs.astral.sh/uv/): 0.8.x
- Bun : 1.2.x

bunをインストールします。

```sh
brew install oven-sh/bun/bun
```

コンテナを起動します。

```sh
# build images
docker compose --profile host build

# run containers
docker compose --profile host up -d

# run db migration
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${DB_URL}'
docker compose run --rm db-migration bash -c 'make migrate-up db_url=${TEST_DB_URL}'
```

parserを実行するには以下のコマンドを実行します。

```sh
cd parser
uv sync --locked
uv run download_and_parse.py --year 2025 --output-path kdb_2025.json
```

backを実行するには以下のコマンドを実行します。

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

frontを実行するには以下のコマンドを実行します。

```sh
cd front
bun install
bun run dev
```

sponsorshipを実行するには以下のコマンドを実行します。

```sh
cd sponsorship
bun install
bun run dev
```

http://localhost:4000 または http://localhost:4000/sponsorship で Twin:te が使用できます。
