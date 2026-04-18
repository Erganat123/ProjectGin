package handlers

import (
    "bookstore-gin/database"
    "bookstore-gin/models"
    "net/http"
    "strconv"
    "time"
    "github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
    userID := c.GetString("username")
    var favs []models.Favorite
    database.DB.Where("user_id = ?", userID).Find(&favs)
    c.JSON(http.StatusOK, favs)
}

func AddFavorite(c *gin.Context) {
    userID := c.GetString("username")
    bookID, _ := strconv.Atoi(c.Param("id"))

    fav := models.Favorite{
        UserID:    userID,
        BookID:    uint(bookID),
        CreatedAt: time.Now(),
    }
    database.DB.Create(&fav)
    c.JSON(http.StatusOK, fav)
}

func RemoveFavorite(c *gin.Context) {
    userID := c.GetString("username")
    bookID := c.Param("id")
    database.DB.Where("user_id = ? AND book_id = ?", userID, bookID).Delete(&models.Favorite{})
    c.Status(http.StatusNoContent)
}