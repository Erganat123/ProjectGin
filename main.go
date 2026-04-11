package main

import (
	"bookstore-gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 🔐 LOGIN (открытый)
	r.POST("/login", handlers.Login)

	// 📖 ОТКРЫТЫЕ ROUTES (можно оставить)
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBook)

	r.GET("/authors", handlers.GetAuthors)
	r.GET("/categories", handlers.GetCategories)

	// 🔒 ЗАЩИЩЕННЫЕ ROUTES
	auth := r.Group("/")
	auth.Use(handlers.AuthMiddleware())

	auth.POST("/books", handlers.CreateBook)
	auth.PUT("/books/:id", handlers.UpdateBook)
	auth.DELETE("/books/:id", handlers.DeleteBook)

	auth.POST("/authors", handlers.CreateAuthor)
	auth.POST("/categories", handlers.CreateCategory)

	r.Run(":8080")
}