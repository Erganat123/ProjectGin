package handlers

import (
	"bookstore-gin/database"
	"bookstore-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	database.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if author.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name required"})
		return
	}

	if err := database.DB.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create author"})
		return
	}

	c.JSON(http.StatusCreated, author)
}