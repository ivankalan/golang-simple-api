package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID			string `json:"id"`
	Title		string `json:"title"`
	Author		string `json:"author"`
	Quantity	int	`json:"quantity"`
}

var books = [] book {
	{ID: "1", Title: "Seni Merayu Tuhan", Author: "Habib Jafar", Quantity: 2},
	{ID: "2", Title: "Filosofi Teras", Author: "Henry Manampiring", Quantity: 5},
	{ID: "3", Title: "Sebuah Seni Untuk Bersikap Bodo Amat", Author: "Mark Manson", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
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
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}