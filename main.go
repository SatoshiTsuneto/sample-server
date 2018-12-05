package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sample-server/goMySql"
)

// POSTリクエストに対する処理
func postJsonHandler(rw http.ResponseWriter, req *http.Request) {
	// ヘッダーの設定
	rw.Header().Set("Content-Type", "application/json")

	// メソッドがPOSTかどうかを確認
	if req.Method != "POST" {
		fmt.Fprint(rw, "Method Not POST.")
		return
	}

	// データの取得
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}

	// データの代入
	var input goMySql.User
	err = json.Unmarshal(body, &input)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}

	// 受け取ったデータを表示
	fmt.Printf("%#v\n", input)

	// クライアントへのレスポンス
	fmt.Fprint(rw, "success post data!")

	// DBにデータを挿入
	goMySql.SqlWrite(input)
}

// GETリクエストに対する処理
func getJsonHandler(rw http.ResponseWriter, req *http.Request) {
	// ヘッダーの設定
	rw.Header().Set("Content-Type", "application/json")

	// DBからデータを取得
	users := goMySql.SqlRead()

	// 構造体をJSONに変換
	response, err := json.Marshal(users)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}

	// クライアントへ取得したデータを投げる
	fmt.Fprint(rw, string(response))
}

func main() {
	// ハンドラの設定
	http.HandleFunc("/post", postJsonHandler)
	http.HandleFunc("/get", getJsonHandler)
	// port9999で受け付ける
	http.ListenAndServe(":9999", nil)
}
