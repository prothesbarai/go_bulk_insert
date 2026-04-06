package main

import (
	"go_bulk_insert/database"
	"go_bulk_insert/logger"
	"go_bulk_insert/routes"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if (err != nil) {
		logger.AppLogger.Error.Println("Couldn't Load env file : ",err)
		os.Exit(1)
	}
}

func main() {
	router := gin.Default()
	database.ConnectDB()
	routes.SetRoutes(router)
	router.Run(":8080")
}