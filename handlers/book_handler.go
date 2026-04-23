package handlers

import (
	"bookstore-gin/database"
	"bookstore-gin/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	db := database.DB

	if cat := c.Query("category"); cat != "" {
		db = db.Where("category_id = ?", cat)
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := 5
	offset := (page - 1) * pageSize

	db.Limit(pageSize).Offset(offset).Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Книга не найдена"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Книга не найдена"})
		return
	}
	c.ShouldBindJSON(&book)
	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Book{}, id)
	c.Status(http.StatusNoContent)
}