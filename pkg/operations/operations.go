package operations

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Quotes_Project/pkg/database"
)

func InsertAuthor(db *sql.DB,author *database.Author) error {
	if author.Name == "" {
		return errors.New("enter the name of Author")
	}
	author.Name = strings.ToLower(author.Name)
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
	name = strings.ToLower(name)
	statement, err :=  db.Prepare("SELECT * FROM author WHERE name LIKE ? ")
	if err != nil {
		return nil, err
	}
	err = statement.QueryRow(name).Scan(&author.Id, &author.Name)
	if err != nil {
		return nil, err
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
	err = statement.QueryRow(id).Scan(&author.Id, &author.Name)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func InsertQuote(db *sql.DB,quote *database.Quote) error {
	quote.AuthorName = strings.ToLower(quote.AuthorName)
	quote.Text = strings.ToLower(quote.AuthorName)
	if quote.AuthorId != 0 {
		author, err := SearchAuthorById(db, quote.AuthorId)
		if err != nil{
			return err
		}
		quote.AuthorId = author.Id
		quote.AuthorName = author.Name
	} else if quote.AuthorName != "" {
		var a *database.Author
		a, err := SearchAuthorByName(db, quote.AuthorName)
		if err != nil {
			a.Name = quote.AuthorName
			err = InsertAuthor(db,a)
			if err != nil {
				return err
			}
		}
		quote.AuthorId = a.Id
		quote.AuthorName = a.Name
	}else if quote.AuthorId == 0 && quote.AuthorName == "" {
		return errors.New("no info about the author")
	}
	var statement, err = db.Exec("INSERT INTO quote(q_text,author_id) VALUES(?, ?)",quote.Text,quote.AuthorId)
	if err != nil {
		return err
	}
	id, err := statement.LastInsertId()
	if id == 0 || err != nil {
		return err
	}
	return nil
}

func FindAuthorQuotes(db *sql.DB, name string) (*database.AuthorQuotes,error) {
	var result database.AuthorQuotes
	if name == "" {
		return nil, errors.New("no author name entered")
	}
	name = strings.ToLower(name)
	author, err := SearchAuthorByName(db, name)
	if err != nil || author.Id == 0 {
		return nil, err
	}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM quote WHERE author_id = %d ",author.Id))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var q database.Quote
		rows.Scan(&q.Id,&q.Text,&q.AuthorId)
		result.Quotes = append(result.Quotes, q)
	}

	return &result, nil
}

func FindQuote(db *sql.DB, quote string) (*database.Quote, error){
		var result database.Quote
		if quote == "" {
			return nil, errors.New("no quote text entered")
		}
		statement, err :=  db.Prepare("SELECT * FROM quote WHERE q_text LIKE ? ")
		if err != nil {
			return nil, err
		}
		quote = strings.ToLower(quote)
		err = statement.QueryRow(quote).Scan(&result.Id, &result.Text,&result.AuthorId)
		if err != nil {
			return nil, err
		}
		author, _ := SearchAuthorById(db,result.AuthorId)
		result.AuthorName = author.Name
		return &result, nil

}