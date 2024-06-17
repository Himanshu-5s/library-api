package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
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

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	if book.Quanity == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not available"})
		return
	}
	book.Quanity--
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	book.Quanity++
	c.IndentedJSON(http.StatusOK, book)
}

func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
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
	router.GET("/books/:id", bookById)
	router.POST("/books", createBooks)
	router.PATCH("/books/:id/checkout", checkoutBook)
	router.PATCH("/books/:id/return", returnBook)
	router.Run("localhost:8080")
}
