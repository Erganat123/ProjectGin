package main

import (
    "bookstore-gin/database"
    "bookstore-gin/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    database.InitDB()

    r := gin.Default()

    r.POST("/login", handlers.Login)
    r.GET("/books", handlers.GetBooks)
    r.GET("/books/:id", handlers.GetBook)
    r.GET("/authors", handlers.GetAuthors)
    r.GET("/categories", handlers.GetCategories)

    auth := r.Group("/")
    auth.Use(handlers.AuthMiddleware())
    {
        auth.POST("/books", handlers.CreateBook)
        auth.PUT("/books/:id", handlers.UpdateBook)
        auth.DELETE("/books/:id", handlers.DeleteBook)
        
        auth.POST("/authors", handlers.CreateAuthor)
        auth.POST("/categories", handlers.CreateCategory)

        auth.GET("/books/favorites", handlers.GetFavorites)
        auth.PUT("/books/:id/favorites", handlers.AddFavorite)
        auth.DELETE("/books/:id/favorites", handlers.RemoveFavorite)
    }

    r.Run(":8080")
}