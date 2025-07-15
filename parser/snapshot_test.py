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

    def setUp(self):
        # スナップショットJSONファイルが保存されているディレクトリ
        self.test_data_dir = os.path.join(os.path.dirname(os.path.abspath(__file__)), "__tests__")
        if not os.path.exists(self.test_data_dir):
            # __tests__ディレクトリが存在しない場合、更新モードでなければエラーとする
            if not TestRunOldSnapshot.update_snapshots:
                 self.fail(f"テストデータディレクトリが見つかりません: {self.test_data_dir}。JSON生成スクリプトを実行するか、--update-snapshots を付けて実行してください。")
            else:
                # 更新モードでディレクトリが存在しない場合、作成する
                os.makedirs(self.test_data_dir, exist_ok=True)


    def test_run_old_snapshots(self):
        """
        run_old関数の出力が2019年から2022年までの保存されたスナップショットと一致することをテストします。
        UPDATE_SNAPSHOTSがTrueの場合、不一致またはファイルが存在しない場合はスナップショットを更新します。
        """
        years = range(2019, 2023) # テスト対象の年を2019年から2022年まで（2023年を含まない）

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
                        print(f"年 {year} のスナップショットを {snapshot_path} で更新しています。")
                        with open(snapshot_path, "w", encoding="utf-8") as f:
                            # 差分を見やすくするためにJSONを整形して保存
                            json.dump(actual_data, f, ensure_ascii=False, indent=4)
                        print(f"年 {year} のスナップショットが更新されました。")
                        # 更新した場合、この年のテストケースは更新モードでは合格とみなす
                        # ここではassertEqualは不要。更新が目的のため
                    else:
                        # 更新モードでなく、不一致またはスナップショットが見つからない場合、テストを失敗させる
                        self.assertEqual(expected_data, actual_data, f"年 {year} のスナップショットの不一致、またはスナップショットが見つかりません (--update-snapshots を使用してください)。")
                else:
                    # データが一致する場合、テストは合格
                    self.assertEqual(expected_data, actual_data, f"年 {year} のスナップショットは既に一致しています。")


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='run_old関数のスナップショットテストを実行します。')
    parser.add_argument('-u', '--update-snapshots', action='store_true',
                        help='不一致がある場合、またはスナップショットが見つからない場合にスナップショットを更新します。')
    # unittest.mainが自身の引数（例: -v, -f）を処理できるようにparse_known_argsを使用
    args, argv_remainder = parser.parse_known_args()

    # パースされた引数に基づいてクラス属性を設定
    TestRunOldSnapshot.update_snapshots = args.update_snapshots

    # 残りの引数をunittest.mainに渡す
    unittest.main(argv=sys.argv[:1] + argv_remainder)
