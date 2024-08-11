#!/usr/bin/env sh

# TODO: mainに変更する
git clone --depth 1 -b feature/staging https://github.com/twin-te/twinte-back /twinte-back
migrate -database ${DB_URL} -path /twinte-back/db/migrations up
