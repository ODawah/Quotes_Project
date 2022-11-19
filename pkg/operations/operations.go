package operations

import (
	"database/sql"
	"errors"
	"fmt"
)

func InsertAuthor(database *sql.DB,name string) error {
	var statement, err = database.Prepare("INSERT INTO Author(name) VALUES(?)")
	if err != nil {
		return err
	}
	statement.Exec(name)
	return nil
}

func SearchAuthor(database *sql.DB, name string) (int, error) {
	if database == nil {
		return 0, errors.New("error connecting to the database")
	}
	name = "%" + name + "%"
	var rows, err = database.Query(fmt.Sprintf("SELECT * FROM author WHERE name LIKE \"%s\" ", name))
	if err != nil {
		return 0, err
	}
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}
	return id, nil
}



