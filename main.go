package main

import (
	"fmt"
	"strconv"

	db "github.com/Quotes_Project/pkg/database"
	"github.com/Quotes_Project/pkg/operations"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	statement, _ := database.Prepare("INSERT INTO quote(q_text,author_id) VALUES(?,?)")
	statement.Exec("isn't it",1)
	s := "omar"
	s = "%" + s + "%"
	var rows, _ = database.Query(fmt.Sprintf("SELECT q_id FROM quote "))
	fmt.Println(rows)
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}

	res, err := operations.SearchAuthor(database ,"omar")

	fmt.Println("the result" + strconv.Itoa(res) + err.Error())
}

