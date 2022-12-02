package handlers

import (
	"net/http"

	db "github.com/Quotes_Project/pkg/database"
	"github.com/Quotes_Project/pkg/operations"
	"github.com/gin-gonic/gin"
)

func GetQuote(c *gin.Context) {
	var input *db.SearchQuote
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't Bind the json body"})
		return
	}
	quote, err := operations.FindQuote(db.DB, input.Text)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quote)

}

func GetAuthorQuotes(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No name inserted"})
		return
	}
	Quotes, err := operations.FindAuthorQuotes(DB, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Quotes)

}

func CreateQuote(c *gin.Context)  {
	var input *db.Quote
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = operations.InsertQuote(db.DB, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)

}