package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Writer string  `json:"writer"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Book One", Writer: "Jhon", Price: 1.0},
	{ID: "2", Title: "Book Two", Writer: "Doe", Price: 2.0},
	{ID: "3", Title: "Book Three", Writer: "NoOne", Price: 3.0},
	{ID: "4", Title: "Book Four", Writer: "SomeOne", Price: 4.0},
	{ID: "5", Title: "Book Five", Writer: "Ghost", Price: 5.0},
}

func main() {
	router := gin.Default()
	router.GET("/api/books", getBooks)
	router.GET("/api/books/:id", getBookByID)
	router.POST("/api/books", postBooks)
	router.DELETE("/api/books/:id", deleteBook)

	router.Run("localhost:8000")
}

// Get All Books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Get Book By ID
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// Add New Book
func postBooks(c *gin.Context) {
	var newBook book

	if error := c.BindJSON(&newBook); error != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// Delete Book
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			c.IndentedJSON(http.StatusOK, books)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

//Update a Book
