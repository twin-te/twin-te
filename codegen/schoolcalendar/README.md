## 概要

`event`と`module_detail`に関するデータの前処理を行います。  
これらのデータは DB に保存されるのではなく JSON として管理されます。

## 手順

`event`を例にコード生成の手順を示します。

1. `./event`配下に年度毎のデータを JSON 形式で配置する
   - これは[筑波大学のウェブサイト](https://www.tsukuba.ac.jp/campuslife/calendar-school/) などを参考に手作業で作成する
2. `./generate.py`を実行する
3. `twin-te/back/module/schoolcalendar/data/event_gen.json`が生成される
