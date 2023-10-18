package routers

import (
	"godev/go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetAllOrders)
	router.GET("/orders/:orderID", controllers.GetOrderByID)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)
	router.DELETE("orders/:orderID", controllers.DeleteOrder)
	return router
}
