package goMySql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string
	Age  int
}

func SqlWrite(user User) {
	// DBに接続　
	db, err := sql.Open("mysql", "ユーザ名:パスワード@/データベース名")
	if err != nil {
		return
	}
	defer db.Close()

	// insert
	stmt, err := db.Prepare("INSERT INTO テーブル名 VALUES (?, ?)")
	if err != nil {
		return
	}
	_, err = stmt.Exec(user.Name, user.Age)
	if err != nil {
		return
	}
	defer stmt.Close()
}

func SqlRead() ([]User) {
	// DBに接続
	db, err := sql.Open("mysql", "ユーザ名:パスワード@/データベース名")
	if err != nil {
		return nil
	}
	defer db.Close()

	// Select
	rows, err := db.Query("SELECT * FROM テーブル名")
	if err != nil {
		return nil
	}
	defer rows.Close()

	// 取得したデータを構造体へマッピング
	var (
		users []User
		user  User
	)
	for rows.Next() {
		rows.Scan(&user.Name, &user.Age)
		users = append(users, user)
	}
	return users
}
