// controllers/bookController.go
package controllers

import (
	"library-api/database"
	"library-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func SearchBooks(c *gin.Context) {
	title := c.Query("title")
	var books []models.Book
	database.DB.Where("title LIKE ?", "%"+title+"%").Find(&books)
	c.JSON(http.StatusOK, books)
}

func BorrowBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	if !book.IsAvailable {
		c.JSON(http.StatusConflict, gin.H{"error": "Buku sedang dipinjam"})
		return
	}

	book.IsAvailable = false
	database.DB.Save(&book)

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dipinjam", "link": "/books/" + id + "/file"})
}

func ReturnBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	book.IsAvailable = true
	database.DB.Save(&book)

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dikembalikan"})
}

func GetBookFile(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	if book.IsAvailable {
		c.JSON(http.StatusForbidden, gin.H{"error": "Buku belum dipinjam"})
		return
	}

	c.File(book.FilePath)
}
