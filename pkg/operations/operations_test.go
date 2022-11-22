package operations

import (
	"database/sql"
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
// Function InsertAuthor Function in Author (Right input)
func TestInsertAuthor(t *testing.T) {
	p := db.Author{Name: "nilson mandella"}
	err := InsertAuthor(data, &p)
	if err != nil {
	t.Fatal(err)
	}
	rows, _ := data.Query("SELECT * FROM author")
	inc := 1
	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		if id != inc {
			t.Fatal(fmt.Sprintf("got = %d :  expected = %d", id, inc))
		}
		inc += 1
	}

}

// Function SearchAuthor Function in Author (Right input)
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
