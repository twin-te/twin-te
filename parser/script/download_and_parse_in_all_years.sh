#!/usr/bin/env bash

cd $(dirname $0) || exit
cd ../

if [ ! -d "./data/parsed" ]; then
  mkdir -p ./data/parsed
fi

for year in {2019..2024}
do
  python download_and_parse.py --year $year --output-path ./data/parsed/$year.json
done
