package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
}

func main() {
	db, err := sql.Open("sqlite3", "./db1.db")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hello")
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY, 
		user_id INTEGER, 
		title TEXT,
		content TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`)

	if err != nil {
		fmt.Println(err)
	}
	statement.Exec()

	statement, err = db.Prepare(`INSERT INTO posts (user_id, title, content) VALUES (?,?,?)`)
	if err != nil {
		fmt.Println(err)
	}
	statement.Exec(1, "test", "lorem ispum hello world")

	rows, err := db.Query("SELECT content FROM posts")
	if err != nil {
		fmt.Println(err)
	}
	var content string
	for rows.Next() {
		rows.Scan(&content)
		defer rows.Close()
		fmt.Println(content)
	}
}
