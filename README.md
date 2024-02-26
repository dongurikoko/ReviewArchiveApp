# ReviewArchiveApp

以下のサイトに制作動機や仕様、設計などをまとめています。(実装途中なので適宜変更中)  
https://kindhearted-roast-aac.notion.site/github-28e72a6593e0407e9c38430cfd8f9ec3?pvs=4

frontend側の構造
App.tsx → index.tsx → index.html → ブラウザ

### install
npm install react-router-dom

### 起動
`docker compose up`

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
export MYSQL_PORT=3307                                            
export MYSQL_DATABASE=review_archive_api

2.backendの通信開始
go run ./cmd/main.go
```
frontendフォルダ下で以下を行う。

react起動
`npm start`)

## 動作
動作はこんな感じ(デザインには目を瞑っていただけると...)


https://github.com/dongurikoko/ReviewArchiveApp/assets/108347471/d176c015-c60f-48ec-8f34-4308ccca51d6


