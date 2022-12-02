package handlers

import (
	"net/http"
	"strconv"

	db "github.com/Quotes_Project/pkg/database"
	"github.com/Quotes_Project/pkg/operations"
	"github.com/gin-gonic/gin"
)

func GetAuthorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author, err := operations.SearchAuthorById(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, author)

}

func GetAuthorByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No name inserted"})
		return
	}
	author, err := operations.SearchAuthorByName(db.DB, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, author)
}

func CreateAuthor(c *gin.Context) {
	var input *db.Author
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = operations.InsertAuthor(db.DB, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)

}
