package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	FilePath    string `json:"file_path"`
	IsAvailable bool   `json:"is_available"`
}
