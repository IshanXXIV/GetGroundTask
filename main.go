package main

import (
	"github.com/IshanXXIV/GetGroundTask/models"
	"github.com/IshanXXIV/GetGroundTask/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/up", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "It is up",
		})
	})

	models.ConnectDB()

	router := gin.Default()
	routes.MapUrls()
	err := router.Run(":8080")

	panic(err)

}
