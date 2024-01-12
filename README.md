# ReviewArchiveApp

以下のサイトに制作動機や仕様、設計などをまとめています。(実装途中なので適宜変更中)  
https://kindhearted-roast-aac.notion.site/github-28e72a6593e0407e9c38430cfd8f9ec3?pvs=4

環境変数の設定(OS毎に調整) 
export MYSQL_USER=root
export MYSQL_PASSWORD=review-archive
export MYSQL_HOST=127.0.0.1
export MYSQL_PORT=3307                                            
export MYSQL_DATABASE=review_archive_api

backendの通信開始
backendフォルダ下でgo run ./cmd/main.go