package main

import (
	"godev/go-rest-api/routers"
)

func main() {
	// r := gin.Default()

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Hello, World!",
	// 	})
	// })

	// r.Run(":8280")
	routers.StartServer().Run(":8382")
}
