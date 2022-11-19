package operations

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	db "github.com/Quotes_Project/pkg/database"
)

// Global DB variable
var database *sql.DB

// Initializing DB
func TestDb(t *testing.T) {
	var err error
	database, err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

// Function InsertAuthor Function in Author (Right input)
func TestInsertAuthor(t *testing.T) {
	err := InsertAuthor(database, "omar")
	if err != nil {
	t.Fatal(err)
	}
	rows, _ := database.Query("SELECT * FROM author")
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
func TestSearchAuthor(t *testing.T) {
	id, err := SearchAuthor(database, "omar")
	if err != nil {
		t.Fatal(err)
	}
	if id == 1 {
		t.Fatal("error not in database")
	}

}
