import unittest
import json
import os
import sys
import argparse

# 現在のスクリプトのディレクトリ（twin-te/parser）を取得
current_script_dir = os.path.dirname(os.path.abspath(__file__))
# モジュール（download_and_parseなど）をインポートできるように、このディレクトリをsys.pathに追加
sys.path.insert(0, current_script_dir)

from download_and_parse import run_old

class TestRunOldSnapshot(unittest.TestCase):
    # スナップショットを更新するかどうかを格納するクラス属性
    update_snapshots = False
    # テスト対象の年を格納するクラス属性
    years_to_test = None

    def setUp(self):
        # スナップショットJSONファイルが保存されているディレクトリ
        self.test_data_dir = os.path.join(os.path.dirname(os.path.abspath(__file__)), "__tests__")
        if not os.path.exists(self.test_data_dir):
            # __tests__ディレクトリが存在しない場合、更新モードでなければエラーとする
            if not TestRunOldSnapshot.update_snapshots:
                 self.fail(f"テストデータディレクトリが見つかりません: {self.test_data_dir}。--update-snapshots を付けて実行してください。")
            else:
                # 更新モードでディレクトリが存在しない場合、作成する
                os.makedirs(self.test_data_dir, exist_ok=True)


    def test_run_old_snapshots(self):
        """
        kdbのパース結果が保存されたスナップショットと一致することをテストします。
        UPDATE_SNAPSHOTSがTrueの場合、不一致またはファイルが存在しない場合はスナップショットを更新します。
        指定された年がある場合、その年のみをテストします。
        """
        if TestRunOldSnapshot.years_to_test:
            years = TestRunOldSnapshot.years_to_test
        else:
            years = range(2019, 2023) # デフォルトのテスト対象年

        for year in years:
            with self.subTest(year=year):
                snapshot_filename = f"{year}.json"
                snapshot_path = os.path.join(self.test_data_dir, snapshot_filename)

                # 現在の出力を取得するために関数を実行
                actual_json_string = run_old(year=year)
                actual_data = json.loads(actual_json_string)

                expected_data = None
                # スナップショットファイルが存在するか確認
                if os.path.exists(snapshot_path):
                    with open(snapshot_path, "r", encoding="utf-8") as f:
                        expected_data = json.load(f)

                # 比較または更新
                if expected_data is None or expected_data != actual_data:
                    if TestRunOldSnapshot.update_snapshots:
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
    parser = argparse.ArgumentParser(description='run_old関数のスナップショットテストを実行します。')
    parser.add_argument('-u', '--update-snapshots', action='store_true',
                        help='不一致がある場合、またはスナップショットが見つからない場合にスナップショットを更新します。')
    parser.add_argument('-y', '--years', type=str,
                        help='テストする年をカンマ区切りで指定します (例: "2019,2020")。')
    # unittest.mainが自身の引数（例: -v, -f）を処理できるようにparse_known_argsを使用
    args, argv_remainder = parser.parse_known_args()

    # パースされた引数に基づいてクラス属性を設定
    TestRunOldSnapshot.update_snapshots = args.update_snapshots
    if args.years:
        TestRunOldSnapshot.years_to_test = [int(y.strip()) for y in args.years.split(',')]

    # 残りの引数をunittest.mainに渡す
    unittest.main(argv=sys.argv[:1] + argv_remainder)
