import unittest
import json
import os
import sys
import argparse
import io

from kdb_parser import parse


class TestParseSnapshot(unittest.TestCase):
    # スナップショットを更新するかどうかを格納するクラス属性
    update_snapshots = False

    def setUp(self):
        # スナップショットJSONファイルが保存されているディレクトリ
        self.test_data_dir = os.path.join(os.path.dirname(os.path.abspath(__file__)), "__tests__")
        if not os.path.exists(self.test_data_dir):
            # __tests__ディレクトリが存在しない場合、更新モードでなければエラーとする
            if not TestParseSnapshot.update_snapshots:
                 self.fail(f"テストデータディレクトリが見つかりません: {self.test_data_dir}。--update-snapshots を付けて実行してください。")
            else:
                # 更新モードでディレクトリが存在しない場合、作成する
                os.makedirs(self.test_data_dir, exist_ok=True)

        # テスト対象のExcelファイルのパス
        self.excel_file_path = os.path.join(self.test_data_dir, "kdb_2021.xlsx")


    def test_parse_2021_snapshot(self):
        """
        kdb.parse関数の結果が保存されたスナップショットと一致することをテストします。
        UPDATE_SNAPSHOTSがTrueの場合、不一致またはファイルが存在しない場合はスナップショットを更新します。
        """
        year = 2021 # テスト対象の年を2021に固定
        snapshot_filename = f"{year}.json"
        snapshot_path = os.path.join(self.test_data_dir, snapshot_filename)

        # 事前にダウンロードされたExcelファイルを読み込む
        if not os.path.exists(self.excel_file_path):
            self.fail(f"テスト用のExcelファイルが見つかりません: {self.excel_file_path}。事前にダウンロードしてください。")

        with open(self.excel_file_path, "rb") as f:
            xlsx_bytes = f.read()

        # parse関数を直接実行し、現在の出力を取得
        actual_data = parse(io.BytesIO(xlsx_bytes))

        expected_data = None
        # スナップショットファイルが存在するか確認
        if os.path.exists(snapshot_path):
            with open(snapshot_path, "r", encoding="utf-8") as f:
                expected_data = json.load(f)

        # 比較または更新
        if expected_data is None or expected_data != actual_data:
            if TestParseSnapshot.update_snapshots:
                print(f"{year} 年のスナップショットを {snapshot_path} で更新しています。")
                with open(snapshot_path, "w", encoding="utf-8") as f:
                    # 差分を見やすくするためにJSONを整形して保存
                    json.dump(actual_data, f, ensure_ascii=False, indent=4)
                print(f"{year} 年のスナップショットが更新されました。")
            else:
                # 更新モードでなく、不一致またはスナップショットが見つからない場合、テストを失敗させる
                self.assertEqual(expected_data, actual_data, f"{year} 年のスナップショットの不一致、またはスナップショットが見つかりません (--update-snapshots を使用してください)。")
        else:
            # データが一致する場合、テストは合格
            self.assertEqual(expected_data, actual_data, f"{year} 年のスナップショットは既に一致しています。")


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='kdb.parse関数のスナップショットテストを実行します。')
    parser.add_argument('-u', '--update-snapshots', action='store_true',
                        help='不一致がある場合、またはスナップショットが見つからない場合にスナップショットを更新します。')

    # unittest.mainが自身の引数（例: -v, -f）を処理できるようにparse_known_argsを使用
    args, argv_remainder = parser.parse_known_args()

    # パースされた引数に基づいてクラス属性を設定
    TestParseSnapshot.update_snapshots = args.update_snapshots

    # 残りの引数をunittest.mainに渡す
    unittest.main(argv=sys.argv[:1] + argv_remainder)
