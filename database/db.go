package database

import (
	"library-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=12345 dbname=librarydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal koneksi database")
	}

	db.AutoMigrate(&models.Book{})
	DB = db
}
