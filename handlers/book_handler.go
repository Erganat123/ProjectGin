package handlers

import (
	"net/http"
	"strconv"

	"bookstore-gin/models"

	"github.com/gin-gonic/gin"
)

var Books []models.Book
var bookID = 1

func GetBooks(c *gin.Context) {
	category := c.Query("category")
	pageStr := c.Query("page")
	page := 1
	if pageStr != "" {
		p, _ := strconv.Atoi(pageStr)
		if p > 0 {
			page = p
		}
	}

	pageSize := 5
	start := (page - 1) * pageSize
	end := start + pageSize

	filtered := []models.Book{}
	for _, b := range Books {
		if category != "" {
			cid, _ := strconv.Atoi(category)
			if b.CategoryID != cid {
				continue
			}
		}
		filtered = append(filtered, b)
	}

	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	c.JSON(http.StatusOK, filtered[start:end])
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if book.Title == "" || book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	book.ID = bookID
	bookID++
	Books = append(Books, book)
	c.JSON(http.StatusCreated, book)
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, b := range Books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, b := range Books {
		if b.ID == id {
			var newBook models.Book
			if err := c.ShouldBindJSON(&newBook); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			newBook.ID = id
			Books[i] = newBook
			c.JSON(http.StatusOK, newBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, b := range Books {
		if b.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
