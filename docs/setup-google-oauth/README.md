# Google OAuth の設定

このドキュメントでは Twin:te の開発環境を立ち上げるために必要な Google OAuth のクライアントID, クライアントシークレットの取得方法を説明します。

## Google Cloud でプロジェクトを作成

<https://developers.google.com/workspace/guides/create-project> に従いプロジェクトを作成します。

## 認証情報の取得

[認証情報](https://console.cloud.google.com/apis/credentials) から [認証情報の作成] > [OAuth クライアント ID] を選択します。  
アプリケーションの種類は「ウェブアプリケーション」を選択し、名前を入力します。  
「承認済みの JavaScript 生成元」には `http://localhost:8080/auth/v3/google/callback` を入力します。

[!image](./client-id.png)

その後、生成されたクライアント ID とクライアントシークレットを入手します。
