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

	createTables()
}

func createTables() {
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			name TEXT,
			email TEXT,
			password TEXT,
			image_url TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Println(err)
		return
	}
	if _, err := statement.Exec(); err != nil {
		log.Println(err)
		return
	}

	statement, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS posts (
		id TEXT PRIMARY KEY, 
		user_id TEXT,
		FOREGIN KEY (user_id) REFERENCES users(id)
		title TEXT,
		content TEXT,
		image_url TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Println(err)
		return
	}
	if _, err := statement.Exec(); err != nil {
		log.Println(err)
		return
	}
}
