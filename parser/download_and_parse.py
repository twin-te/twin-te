import argparse
import io
import json
import pathlib
import sentry_sdk
import os

from kdb_downloader import KDBDownloader
from kdb_parser import parse
from kdb_parser_after_hook import parse_after_hook


def run_old(year: int) -> str:
    kdb_downloader = KDBDownloader()
    xlsx_bytes = kdb_downloader.download_excel(year=year)

    course_data = parse(io.BytesIO(xlsx_bytes))

    json_data = json.dumps(course_data, ensure_ascii=False, separators=(",", ":"))

    json_data = json_data.replace("\\r\\n", "\\n")
    json_data = json_data.replace("\\n", "\\r\\n")

    return json_data


def run_all(year: int) -> str:
    return parse_after_hook(run_old(year=year))


def main():
    dsn = os.getenv("PARSER_SENTRY_DSN")
    if dsn:
        sentry_sdk.init(dsn=dsn, send_default_pii=True)
    parser = argparse.ArgumentParser()

    parser.add_argument("--year", type=int, required=True, help="academic year")
    parser.add_argument("--output-path", type=str, required=True)

    args = parser.parse_args()

    data = run_all(args.year)

    output_path = pathlib.Path(args.output_path)
    with output_path.open("w") as f:
        f.write(data)


if __name__ == "__main__":
    main()
