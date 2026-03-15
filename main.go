package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/look4suman/events-api/db"
	"github.com/look4suman/events-api/routes"
)

func main() {

	db.InitDB()

	// Create a Gin router with default middleware (logger and recovery)
	server := gin.Default()

	routes.RegisterRoutes(server)

	// Start server on port 8081 (default - 8080)
	if err := server.Run(":8081"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
