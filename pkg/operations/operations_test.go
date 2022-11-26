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
		input  db.Author
		output error
	}
	tests := []test{
		{input: db.Author{Name: "omar"}, output: nil},
		{input: db.Author{Name: ""}, output: errors.New("enter the name of Author")},
	}

	for _, tc := range tests {
		got := InsertAuthor(data, &tc.input)

		if fmt.Sprint(got) != fmt.Sprint(tc.output) {
			t.Fatal(fmt.Sprintf("expected: %s, got: %s", tc.output, got))
		}
	}

}

func TestSearchAuthorByName(t *testing.T) {
	type test struct {
		input  string
		outErr error
	}

	tests := []test{
		{input: "omar", outErr: nil},
		{input: "maged", outErr: errors.New("sql: no rows in result set")},
		{input: "", outErr: errors.New("no name entered")},
	}

	for _, tc := range tests {
		got, gotErr := SearchAuthorByName(data, tc.input)
		if fmt.Sprint(gotErr) != fmt.Sprint(tc.outErr) && got.Id != 0 {
			t.Fatal(fmt.Sprintf("expected: %s, got: %s", tc.outErr, gotErr))
		}
	}

}

func TestSearchAuthorById(t *testing.T) {
	type test struct {
		input  int
		output *db.Author
	}

	tests := []test{
		{input: 1, output: &db.Author{Name: "omar", Id: 1}},
		{input: 50, output: nil},
		{input: 0},
	}

	for _, tc := range tests {
		got, gotErr := SearchAuthorById(data, tc.input)
		if got != tc.output && gotErr != nil {
			t.Fatal(fmt.Sprintf("expected: %s, got: %s", tc.output.Name, gotErr))
		}
	}

}

func TestInsertQuote(t *testing.T) {
	type test struct {
		input  db.Quote
		output error
	}

	tests := []test{
		{input: db.Quote{Text: "work hard3.6", AuthorName: "omar"}, output: nil}, // insert Quote with author in db
		{input: db.Quote{Text: "bring the light3.6", AuthorName: "badr"}, output: nil}, // insert Quote with author not in db
		{input: db.Quote{Text: "build your career3.6", AuthorId: 2}, output: nil}, // insert Quote with author id in db
		{input: db.Quote{Text: "build your path3.6", AuthorId: 100}, output: errors.New("sql: no rows in result set")}, // insert Quote with author id not in db
		{input: db.Quote{Text: "look at the sun3.6"}, output: errors.New("no info about the author")}, // insert Quote with no author info
	}

	for i, tc := range tests {
		gotErr := InsertQuote(data, &tc.input)
		if fmt.Sprint(tc.output) != fmt.Sprint(gotErr) {
			t.Fatal(fmt.Sprintf("%d: expected: %s, got: %s", i, fmt.Sprint(tc.output), fmt.Sprint(gotErr)))
		}
	}
}


func TestFindAuthorQuotes(t *testing.T) {
	type test struct {
		input  string
		output error
	}

	tests := []test{
		{input: "omar", output: nil},
		{input: "peter", output: errors.New("sql: no rows in result set") },
		{input: "", output: errors.New("no author name entered")},
	}

	for i, tc := range tests {
		got ,gotErr := FindAuthorQuotes(data, tc.input)
		if got == nil && fmt.Sprint(gotErr) !=  fmt.Sprint(tc.output) {
			t.Fatal(fmt.Sprintf("%d: expected: %s, got: %s", i, fmt.Sprint(tc.output), fmt.Sprint(gotErr)))
		}
	}

}

func TestFindQuote(t *testing.T) {
	type test struct {
		input  string
		output error
	}

	tests := []test{
		{input: "work hard", output: nil},
		{input: "not in db quote", output: errors.New("sql: no rows in result set") },
		{input: "", output: errors.New("no quote text entered")},
	}

	for i, tc := range tests {
		got ,gotErr := FindQuote(data, tc.input)
		t.Log(got)
		if got == nil && fmt.Sprint(gotErr) !=  fmt.Sprint(tc.output) {
			t.Fatal(fmt.Sprintf("%d: expected: %s, got: %s", i, fmt.Sprint(tc.output), fmt.Sprint(gotErr)))
		}
	}

}




