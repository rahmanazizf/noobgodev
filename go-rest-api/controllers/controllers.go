package controllers

import (
	"fmt"
	"godev/go-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var OrderData = []models.Order{}

var connectDB *models.Connection

func NewConnection(conn *gorm.DB) *models.Connection {
	return &models.Connection{
		DB: conn,
	}
}

func EstablishConnection(conn *models.Connection) {
	connectDB = conn
}

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Binding JSON to model failed",
		})
		return
	}

	// store data to database
	err = connectDB.CreateRec(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    newOrder,
		"message": "success",
	})
}

var found = false
var orderData *models.Order

func GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("orderID")
	for _, o := range OrderData {
		if id, _ := strconv.Atoi(id); o.OrderID == id {
			found = true
			orderData = &o
		}
	}
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    orderData,
			"message": fmt.Sprintf("Order ID %s is not available", id),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    orderData,
		"message": "sucess",
	})
}

func GetAllOrders(ctx *gin.Context) {
	var orders []models.Order
	err := connectDB.GetAllRec(orders)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": "success",
	})
}

// update order dengan fungsi ini dilakukan dengan cara menghapus semua item terkait order yang ditargetkan
func UpdateOrder(ctx *gin.Context) {
	var existingOrders models.Order
	id, _ := strconv.Atoi(ctx.Param("orderID"))
	var updatedOrder models.Order
	ctx.ShouldBindJSON(&updatedOrder)
	err := connectDB.UpdateRecByID(&existingOrders, &updatedOrder, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Order ID not found!",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    existingOrders,
		"message": "update success",
	})

}

func DeleteOrder(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("orderID"))

	var orders []models.Order
	err := connectDB.DeleteRecByID(&orders, orderID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Order ID not found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": fmt.Sprintf("Deleted an order with id %d", orderID),
	})
}
