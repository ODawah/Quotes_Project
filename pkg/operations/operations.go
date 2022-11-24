package operations

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Quotes_Project/pkg/database"
)

func InsertAuthor(db *sql.DB,author *database.Author) error {
	if author.Name == "" {
		return errors.New("enter the name of Author")
	}
	var statement, err = db.Exec("INSERT INTO Author(name) VALUES(?)", author.Name)
	if err != nil {
		return err
	}
	id, err := statement.LastInsertId()
	if err != nil || id == 0 {
		return err
	}
	return nil
}

func SearchAuthorByName(db *sql.DB, name string) (*database.Author, error) {
	var author database.Author
	if name == "" {
		return nil, errors.New("no name entered")
	}

	name = "%" + name + "%"
	var rows, err = db.Query(fmt.Sprintf("SELECT * FROM author WHERE name LIKE \"%s\" LIMIT 1", name))
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	for rows.Next() {
		rows.Scan(&author.Id,&author.Name)
	}
	return &author, nil
}

func SearchAuthorById(db *sql.DB, id int) (*database.Author, error) {
	var author database.Author
	if id == 0 {
		return nil, errors.New("no id entered")
	}
	statement, err :=  db.Prepare("SELECT * FROM author WHERE id = ? ")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	err = statement.QueryRow(id).Scan(&author.Id, &author.Name)
	return &author, nil
}

func InsertQuote(database *sql.DB,quote *database.Quote) error {
	if quote.Id != 0 {
		author, err := SearchAuthorById(database, quote.Id)
		if err != nil {
			return err
		}
		quote.AuthorId = author.Id
		quote.AuthorName = author.Name
	} else if quote.AuthorName != "" {
		author, err := SearchAuthorByName(database, quote.AuthorName)
		if err != nil {
			return err
		}
		quote.AuthorId = author.Id
		quote.AuthorName = author.Name
	}
	var statement, err = database.Prepare("INSERT INTO quote(q_text,author_id) VALUES(?, ?)")
	if err != nil {
		return err
	}
	statement.Exec(quote.Text, quote.AuthorId)
	return nil
}