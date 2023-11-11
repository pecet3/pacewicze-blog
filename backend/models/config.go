package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Config() {
	var err error
	db, err = sql.Open("sqlite3", "./db1.db")
	if err != nil {
		log.Println(err)
		return
	}

	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY, 
		user_id INTEGER, 
		title TEXT,
		content TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Println(err)
		return
	}
	statement.Exec()
}
