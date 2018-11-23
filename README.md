## Golangであれこれするサーバアプリ

### プログラム内容
- POST：JSONデータを受け取ったらデータベースに保存
- GET ：データベースからデータを取得して投げる

### セットアップ
- go get "github.com/go-sql-driver/mysql"

### テスト
- curl -H "Accept: application/json" -H "Content-type: application/json" -X POST http://localhost:9999/post -d '{"Name": "sample", "Age": 9999}'
- curl -X GET http://localhost:9999/get