package routers

import (
	"godev/go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/orders", controllers.CreateOrder)
	return router
}
