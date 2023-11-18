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
		log.Println("config error:", err)
		return
	}

	err = createTables()
	if err != nil {
		log.Println("config error:", err)
		return
	}

}

func createTables() error {
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			name TEXT,
			email TEXT,
			password TEXT,
			image_url TEXT,
			salt TEXT,
			is_active BOOLEAN DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		return err
	}
	if _, err := statement.Exec(); err != nil {
		return err
	}

	statement, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS posts (
			id TEXT PRIMARY KEY, 
			user_id TEXT,
			title TEXT,
			content TEXT,
			image_url TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {

		return err
	}
	if _, err := statement.Exec(); err != nil {

		return err
	}

	return nil
}
