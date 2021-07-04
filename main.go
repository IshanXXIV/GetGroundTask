package main

import (
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

	router := gin.Default()
	routes.MapUrls(router)
	err := router.Run(":8080")

	panic(err)

}
