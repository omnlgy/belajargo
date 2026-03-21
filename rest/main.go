package main

import (
	"example.com/rest/db"
	"example.com/rest/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()

	server := gin.Default()

	router.RegisterRoutes(server)

	server.Run(":8080")
}
