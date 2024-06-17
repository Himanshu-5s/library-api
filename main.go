package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id      string
	Title   string
	Author  string
	Quanity int
}

var books = []book{
	{Id: "1", Title: "The Alchemist", Author: "Paulo Coelho", Quanity: 10},
	{Id: "2", Title: "The Monk Who Sold His Ferrari", Author: "Robin Sharma", Quanity: 5},
	{Id: "3", Title: "The Power of Now", Author: "Eckhart Tolle", Quanity: 7},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
