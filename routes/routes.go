package routes

import (
	"library-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/search", controllers.SearchBooks)
	r.POST("/books/:id/borrow", controllers.BorrowBook)
	r.POST("/books/:id/return", controllers.ReturnBook)
	r.GET("/books/:id/file", controllers.GetBookFile)
}
