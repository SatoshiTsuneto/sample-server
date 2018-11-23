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
	var response string
	defer func() {
		_, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}
		rw.Header().Set("Content-Type", "application/json")
	}()

	// メソッドがPOST確認
	if req.Method != "POST" {
		response = "Method Not POST."
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response = err.Error()
		fmt.Println(err.Error())
		return
	}

	// データの取得
	var input goMySql.User
	err = json.Unmarshal(body, &input)
	if err != nil {
		response = err.Error()
		fmt.Println(err.Error())
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
	// DBからデータを取得
	users := goMySql.SqlRead()

	// 構造体をJSONに変換
	response, _ := json.Marshal(users)

	// クライアントへ取得したデータを投げる
	defer func() {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(response))
	}()
}

func main() {
	// ハンドラの設定
	http.HandleFunc("/post", postJsonHandler)
	http.HandleFunc("/get", getJsonHandler)
	// port8080で受け付ける
	http.ListenAndServe(":8080", nil)
}
