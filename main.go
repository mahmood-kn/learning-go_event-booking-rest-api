package main

import (
	"exmaple.com/event-booking-rest-api/db"
	"exmaple.com/event-booking-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
