package main

import (
	"library-api/database"
	"library-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Static("/assets", "./assets")
	database.InitDB()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
