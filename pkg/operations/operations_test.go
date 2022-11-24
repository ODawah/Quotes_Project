package operations

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"

	db "github.com/Quotes_Project/pkg/database"
)

// Global DB variable
var data *sql.DB

// Initializing DB
func TestDb(t *testing.T) {
	var err error
	data, err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
}


func TestInsertAuthor(t *testing.T) {
	type test struct {
		input db.Author
		output  error
	}
	err := errors.New("enter the name of Author")
	tests := []test{
		{input: db.Author{Name: "omar"}, output: nil},
		{input: db.Author{Name: ""}, output: err},
	}

	for _, tc := range tests {
		got := InsertAuthor(data, &tc.input)

		if fmt.Sprint(got) != fmt.Sprint(tc.output) {
			t.Fatal(fmt.Sprintf("expected: %s, got: %s", tc.output, got))
		}
	}

}

func TestSearchAuthorByName(t *testing.T) {

	name := "omar"

	author, err := SearchAuthorByName(data, name)
	if err != nil {
		t.Fatal(err)
	}
	if author.Id == 0 {
		t.Log(author.Name)
		t.Log(author.Id)
		t.Fatal("author not in database")
	}

}


func TestSearchAuthorById(t *testing.T) {
	author, err := SearchAuthorById(data, 5)
	if err != nil {
		t.Fatal(err)
	}
	if author.Name == "" {
		t.Log("here")
		t.Fatal("error not in database")
	}

}

func TestInsertQuote_By_Name(t *testing.T) {
	p := db.Quote{
		Text:       "heaven is for real",
		AuthorName: "omar",
		AuthorId:   0,
	}
	err := InsertQuote(data, &p)
	if err != nil {
		t.Fatal(err)
	}
	rows, _ := data.Query("SELECT * FROM quote")
	var got db.Quote

	for rows.Next() {
		rows.Scan(&got.Id, &got.Text, &got.AuthorId)
		if got.Text != p.Text {
			t.Fatal("error")
		}
	}

}

func TestInsertQuote_By_id(t *testing.T) {
	p := db.Quote{
		Id: 2,
		Text:       "heaven is for real",
		AuthorId:   0,
	}
	err := InsertQuote(data, &p)
	if err != nil {
		t.Fatal(err)
	}
	rows, _ := data.Query("SELECT * FROM quote")
	var got db.Quote

	for rows.Next() {
		rows.Scan(&got.Id, &got.Text, &got.AuthorId)
		if got.Text != p.Text {
			t.Fatal("error")
		}
	}

}

