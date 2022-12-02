package database

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once
var DB *sql.DB

// Connect is Function that connects to database
func Connect() (*sql.DB, error) {
	once.Do(func() {
		database, _ := sql.Open("sqlite3", "./quotes.db")
		DB = database
		statement, _ := DB.Prepare("CREATE TABLE IF NOT EXISTS  author (id INTEGER PRIMARY KEY AUTOINCREMENT , name TEXT NOT NULL)")
		statement.Exec()
		statement, _ = DB.Prepare("CREATE TABLE IF NOT EXISTS  quote (q_id INTEGER PRIMARY KEY AUTOINCREMENT , q_text TEXT NOT NULL UNIQUE , author_id INTEGER NOT NULL, FOREIGN KEY (author_id) REFERENCES Author(id))")
		statement.Exec()

	})

	return DB, nil
}
