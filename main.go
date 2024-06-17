package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID      string
	Title   string
	Author  string
	Quanity int
}

var books = []book{
	{ID: "1", Title: "The Alchemist", Author: "Paulo Coelho", Quanity: 10},
	{ID: "2", Title: "The Monk Who Sold His Ferrari", Author: "Robin Sharma", Quanity: 5},
	{ID: "3", Title: "The Power of Now", Author: "Eckhart Tolle", Quanity: 7},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBooks)
	router.Run("localhost:8080")
}
