import argparse

import requests


class KDBDownloader:
    def _download_excel(self, year) -> bytes:
        headers = {
            "Accept-Language": "ja,ja-JP;q=0.9,en;q=0.8",
        }
        data = {
            "pageId": "SB0070",
            "action": "downloadList",
            "hdnFy": year,
            "hdnTermCode": "",
            "hdnDayCode": "",
            "hdnPeriodCode": "",
            "hdnAgentName": "",
            "hdnOrg": "",
            "hdnIsManager": "",
            "hdnReq": "",
            "hdnFac": "",
            "hdnDepth": "",
            "hdnChkSyllabi": "false",
            "hdnChkAuditor": "false",
            "hdnChkExchangeStudent": "false",
            "hdnChkConductedInEnglish": "false",
            "hdnCourse": "",
            "hdnKeywords": "",
            "hdnFullname": "",
            "hdnDispDay": "",
            "hdnDispPeriod": "",
            "hdnOrgName": "",
            "hdnReqName": "",
            "cmbDwldtype": "excel",
        }
        url = "https://kdb.tsukuba.ac.jp/"
        response = requests.post(url, data=data, headers=headers)
        if response.status_code != 200:
            raise Exception("invalid status code")

        return response.content

    def download_excel(self, year=int) -> bytes:
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
