package controllers

import (
	"godev/go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var OrderData = []models.Order{}

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	newOrder.OrderID = len(OrderData) + 1
	OrderData = append(OrderData, newOrder)

	ctx.JSON(http.StatusOK, gin.H{
		"data":    newOrder,
		"message": "success",
	})
}
