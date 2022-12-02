package main

import (
	"log"

	db "github.com/Quotes_Project/pkg/database"
	"github.com/Quotes_Project/pkg/handlers"
	"github.com/Quotes_Project/pkg/operations"
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	_ = operations.InsertAuthor(database, &db.Author{Name: "omar"})
	router.GET("/find_author/id/:id", handlers.GetAuthorByID)
	router.GET("/find_author/name/:name", handlers.GetAuthorByName)
	router.POST("/create_author", handlers.CreateAuthor)

	log.Fatalln(router.Run(":8080"))
}
