#!/usr/bin/env sh

current_year=$(date +%Y)
current_month=$(date +%m)

if [ "$current_month" -ge "04" ]; then
    academic_year=$current_year
else
    academic_year=$((current_year - 1))
fi

tmp_dir="/tmp/twinte-parser"

mkdir -p /tmp/twinte-parser
docker build -t twin-te/twinte-parser --no-cache . 
docker run \
  -v $tmp_dir:$tmp_dir \
  twin-te/twinte-parser \
  python3 /twinte-back/twinte-parser/download_and_parse.py --year $academic_year --output-path $tmp_dir/kdb.json

docker run \
  -v $tmp_dir:$tmp_dir \
  --add-host db:10.0.0.3 \
  --env-file ../.twinte-back.env \
  ghcr.io/twin-te/twinte-back:stg \
  /app update-courses-based-on-kdb --year $academic_year --kdb-json-file-path $tmp_dir/kdb.json
