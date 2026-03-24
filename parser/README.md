# twinte-parser
KdBから授業情報を取得するためのdownloaderとparserです。

## Get Started

```sh
uv sync --locked

uv run download_and_parse.py --year 2023 --output-path output.json
```

## スナップショットテスト

```sh
uv run snapshot_test.py
uv run snapshot_test.py -u # スナップショットテストを更新
```
