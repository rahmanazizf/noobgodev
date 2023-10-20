package controllers

import (
	"fmt"
	"godev/go-rest-api/helpers"
	"godev/go-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// store data to database
	db := connectDB.DB
	db.Create(&newOrder)

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
	db := connectDB.DB
	db.Preload("Items").Find(&orders)
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
	db := connectDB.DB
	res := db.Preload("Items").Where("order_id = ?", id).First(&existingOrders)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    existingOrders,
			"message": fmt.Sprintf("Order ID %d is not available", id),
		})
		return
	}
	// update data
	// delete existing items with related order_id
	var deletedItem []models.Item
	db.Clauses(clause.Returning{}).Where("order_id = ?", existingOrders.OrderID).Delete(&deletedItem)

	existingOrders.CustomerName = updatedOrder.CustomerName
	existingOrders.OrderedAt = updatedOrder.OrderedAt
	existingOrders.Items = updatedOrder.Items
	db.Save(&existingOrders)
	ctx.JSON(http.StatusOK, gin.H{
		"data":    existingOrders,
		"message": "update success",
	})

}

func DeleteOrder(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("orderID"))

	db := connectDB.DB
	var orders []models.Order
	res := db.Clauses(clause.Returning{}).Where("order_id = ?", orderID).Delete(&orders)
	helpers.CheckError(res.Error)

	ctx.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": fmt.Sprintf("Deleted an order with id %d", orderID),
	})
}
