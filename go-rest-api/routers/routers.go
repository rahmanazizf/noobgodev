package routers

import (
	"godev/mid/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/orders", controllers.CreateOrder)
	return router
}
