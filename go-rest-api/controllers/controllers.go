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
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	// generate ID (kalau udah nyambung ke database bakal dihapus)
	if len(OrderData) == 0 {
		newOrder.OrderID = 1
	} else {
		newOrder.OrderID = OrderData[len(OrderData)-1].OrderID + 1
	}
	OrderData = append(OrderData, newOrder)

	// store data to database
	db := connectDB.DB
	res := db.Create(&newOrder)
	fmt.Println(res.RowsAffected)

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
	ctx.JSON(http.StatusOK, gin.H{
		"data":    OrderData,
		"message": "success",
	})
}

func UpdateOrder(ctx *gin.Context) {
	var existingOrder *models.Order
	id := ctx.Param("orderID")
	var updatedOrder *models.Order
	ctx.ShouldBindJSON(&updatedOrder)
	for _, o := range OrderData {
		if id, _ := strconv.Atoi(id); o.OrderID == id {
			found = true
			existingOrder = &o
		}
	}
	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"data":    orderData,
			"message": fmt.Sprintf("Order ID %s is not available", id),
		})
	}

	updatedOrder.OrderID = existingOrder.OrderID
	OrderData = append(RemoveIndex(OrderData, updatedOrder.OrderID), *updatedOrder)
	ctx.JSON(http.StatusOK, gin.H{
		"data":    updatedOrder,
		"message": "update success",
	})

}

func RemoveIndex(slc []models.Order, orderID int) []models.Order {
	var iStop int = -1
	newOrderData := []models.Order{}
	for i, o := range slc {
		if o.OrderID == orderID {
			iStop = i
			break
		}
		newOrderData = append(newOrderData, o)
	}
	if iStop+1 <= len(slc) {
		return append(newOrderData, slc[iStop+1:]...)
	}
	return newOrderData
}

func DeleteOrder(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("orderID"))
	OrderData = RemoveIndex(OrderData, orderID)
	ctx.JSON(http.StatusOK, gin.H{
		"data":    OrderData,
		"message": fmt.Sprintf("Deleted an order with id %d", orderID),
	})
}
