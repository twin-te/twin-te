import argparse

import requests
from bs4 import BeautifulSoup


class KDBDownloader:
    _flowExecutionKey: str
    _session: requests.Session

    def _grant_session(self) -> None:
        url = "https://kdb.tsukuba.ac.jp/"
        response = self._session.get(url)
        if response.status_code != 200:
            raise Exception("invalid status code")

        html = response.text
        soup = BeautifulSoup(html, "html.parser")
        tag = soup.find("input", attrs={"name": "_flowExecutionKey"})
        self._flowExecutionKey = tag.attrs["value"]

    def _search_all(self, year: int) -> None:
        params = {
            "_flowExecutionKey": self._flowExecutionKey,
            "_eventId": "searchOpeningCourse",
            "index": "",
            "locale": "",
            "nendo": year,
            "termCode": "",
            "dayCode": "",
            "periodCode": "",
            "campusCode": "",
            "hierarchy1": "",
            "hierarchy2": "",
            "hierarchy3": "",
            "hierarchy4": "",
            "hierarchy5": "",
            "freeWord": "",
            "_gaiyoFlg": 1,
            "_risyuFlg": 1,
            "_excludeFukaikoFlg": 1,
            "outputFormat": 0,
        }

        url = "https://kdb.tsukuba.ac.jp/campusweb/campussquare.do"
        headers = {"Content-Type": "application/x-www-form-urlencoded"}
        response = self._session.post(url, params=params, headers=headers)
        if response.status_code != 200:
            raise Exception("invalid status code")

        html = response.text
        soup = BeautifulSoup(html, "html.parser")
        tag = soup.find("input", attrs={"name": "_flowExecutionKey"})
        self._flowExecutionKey = tag.attrs["value"]

    def _download_excel(self, year) -> bytes:
        params = {
            "_flowExecutionKey": self._flowExecutionKey,
            "_eventId": "outputOpeningCourseExcel",
            "index": "",
            "locale": "",
            "nendo": year,
            "termCode": "",
            "dayCode": "",
            "periodCode": "",
            "campusCode": "",
            "hierarchy1": "",
            "hierarchy2": "",
            "hierarchy3": "",
            "hierarchy4": "",
            "hierarchy5": "",
            "freeWord": "",
            "_gaiyoFlg": 1,
            "_risyuFlg": 1,
            "_excludeFukaikoFlg": 1,
            "outputFormat": 1,
        }
        url = "https://kdb.tsukuba.ac.jp/campusweb/campussquare.do"
        response = self._session.post(url, params=params)
        if response.status_code != 200:
            raise Exception("invalid status code")

        return response.content

    def download_excel(self, year=int) -> bytes:
        self._session = requests.Session()

        self._grant_session()
        self._search_all(year=year)
        kdb_bytes = self._download_excel(year=year)

        return kdb_bytes


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--year", type=int, required=True, help="academic year")
    parser.add_argument(
        "--output-path",
        type=str,
        required=True,
        help="excel file path such as kdb.xlsx",
    )
    args = parser.parse_args()

    kdb_downloader = KDBDownloader()
    xlsx_bytes = kdb_downloader.download_excel(year=args.year)

    with open(args.output_path, "wb") as f:
        f.write(xlsx_bytes)


if __name__ == "__main__":
    main()
