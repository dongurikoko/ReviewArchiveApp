# ReviewArchiveApp

以下のサイトに制作動機や仕様、設計など諸々をまとめています。(修正途中なので適宜変更中)  
https://kindhearted-roast-aac.notion.site/github-28e72a6593e0407e9c38430cfd8f9ec3?pvs=4

テーブル設計：

<img width="450" alt="スクリーンショット 2024-06-05 2 59 30" src="https://github.com/dongurikoko/ReviewArchiveApp/assets/108347471/db720373-9f5b-45d9-8f90-63c9408e9f13">


### 起動
`docker-compose up`

ブラウザで"http://localhost:3000/"
にアクセス

(
ちなみにdockerを使わない場合：

backendフォルダ下で以下を行う。
```
1.環境変数の設定(OS毎に調整)
export MYSQL_USER=root
export MYSQL_PASSWORD=review-archive
export MYSQL_HOST=127.0.0.1
export MYSQL_PORT=3306                                           
export MYSQL_DATABASE=review_archive_api

2.backendの通信開始
go run ./cmd/main.go
```
frontendフォルダ下で以下を行う。

react起動
`npm start`)

## Firebase Admin SDKのセットアップ  
- Firebaseコンソールからプロジェクトを作成し、プロジェクト設定でサービスアカウントを生成して秘密鍵（JSONファイル）をダウンロード  
- 秘密鍵ファイルのパスを環境変数に設定
ここでは"CREDENTIALS"とする。  
`export CREDENTIALS=/path/to/your/firebase-service-account-file.json`

## 動作
動作はこんな感じ(デザインには目を瞑っていただけると...)


https://github.com/dongurikoko/ReviewArchiveApp/assets/108347471/d176c015-c60f-48ec-8f34-4308ccca51d6


