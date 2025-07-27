#!/usr/bin/env bash

cd $(dirname $0) || exit
cd ../

if [ ! -d "./data/raw" ]; then
  mkdir -p ./data/raw
fi

for year in {2019..2024}
do
  uv run kdb_downloader.py --year $year --output-path ./data/raw/$year.xlsx
done
