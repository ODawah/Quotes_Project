package main

import (
	"log"

	"github.com/Quotes_Project/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()

	router.GET("/find_author/id/:id", handlers.GetAuthorByID)
	router.GET("/find_author/name/:name", handlers.GetAuthorByName)
	router.POST("/create_author", handlers.CreateAuthor)
	router.GET("/find_quote", handlers.GetQuote)
	router.GET("/find_quotes/:name", handlers.GetAuthorQuotes)
	router.POST("/create_quote", handlers.CreateQuote)

	log.Fatalln(router.Run(":8080"))
}
