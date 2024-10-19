# V3 -> V4 のデータ移行スクリプト

このスクリプトは V3 から V4 へのデータ移行を行うためのスクリプトです。

※現状は本番ではなく local の docker container で本番データを入れることを想定しています。

## 使い方

1. V3 DB から CSV (with HEADER, null as 'NULL') を出力します。
2. `/data/raw` に CSV を配置します。
3. トップディレクトリで `docker-compose up db` を実行しておきます。
4. 別のシェルを立ち上げて、この README.md と同じディレクトリで `bash import.sh` を実行します。
