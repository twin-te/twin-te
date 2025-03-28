## 概要

`event`と`module_detail`に関するデータの前処理を行います。  
これらのデータは DB に保存されるのではなく JSON として管理されます。

## 手順

`event`を例にコード生成の手順を示します。

1. `./event`配下に年度毎のデータを JSON 形式で配置する
   - これは[筑波大学のウェブサイト](https://www.tsukuba.ac.jp/campuslife/calendar-school/) などを参考に手作業で作成する
2. `python ./verify.py {年度}` を実行して JSON ファイルを検証する
3. `python ./generate.py`を実行する
4. `twin-te/back/module/schoolcalendar/data/event_gen.json`が生成される
