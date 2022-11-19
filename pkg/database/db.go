package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Connect is Function that connects to database
func Connect() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "./quotes.db")
	if err != nil {
		return nil, err
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS  author (id INTEGER PRIMARY KEY AUTOINCREMENT , name TEXT NOT NULL)")
	if err != nil {
		return nil,err
	}
	statement.Exec()
	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS  quote (q_id INTEGER PRIMARY KEY AUTOINCREMENT , q_text TEXT, author_id INTEGER NOT NULL, FOREIGN KEY (author_id) REFERENCES Author(id))")
	if err != nil {
		return nil,err
	}
	statement.Exec()

	return database, nil
}
