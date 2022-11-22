package controllers

import (
	"assigment-2/config"
	"assigment-2/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CreateOrder(ctx *gin.Context) {
	DB := config.GetDB()
	var order models.Order
	var errTx error
	tx := DB.Begin()
	decoder := json.NewDecoder(ctx.Request.Body)
	order = models.Order{
		OrderedAt: time.Now(),
	}
	err := decoder.Decode(&order)

	if err != nil {
		panic(err)
	}

	if err := tx.Create(&order).Error; err != nil {
		errTx = err
	}

	if errTx != nil {
		ctx.JSON(400, gin.H{
			"message": "Failed Create",
			"err":     errTx,
		})
		tx.Rollback()
	} else {

		ctx.JSON(http.StatusCreated, gin.H{
			"customerName": order.CustomerName,
			"items":        order.Items,
			"orderedAt":    order.OrderedAt,
		})
	}

	tx.Commit()

}

func GetOrders(ctx *gin.Context) {
	orders := []models.Order{}
	config.GetDB().Preload("Items").Find(&orders)
	ctx.JSON(200, gin.H{
		"orders": orders,
	})
}

func GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")
	order := models.Order{}

	if err := config.GetDB().Preload("Items").Where("orders.id = ?", id).First(&order).Error; err != nil {
		ctx.JSON(404, gin.H{
			"message": "Not Found",
			"status":  "Failed",
		})
		ctx.Abort()
		return
	}

	var responseItem []map[string]interface{}

	for i := 0; i < len(order.Items); i++ {
		responseItem = []map[string]interface{}{
			{
				"itemCode":    order.Items[i].ItemCode,
				"description": order.Items[i].Description,
				"quantity":    order.Items[i].Quantity,
			},
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Found Product",
		"order": map[string]interface{}{
			"customerName": order.CustomerName,
			"orderedAt":    order.OrderedAt,
			"items":        responseItem,
		},
	})

}

func UpdateOrder(ctx *gin.Context) {
	db := config.GetDB()
	orderId := ctx.Param("orderId")
	order := models.Order{}
	err := db.First(&order, "id = ?", orderId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"item_code", "description", "quantity"}),
	}).Where("orderid = ?", orderId).Create(&order.Items)

	ctx.JSON(200, gin.H{
		"data": order,
	})
}

func DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	order := models.Order{}

	config.GetDB().Delete(&order, orderId)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success deleted order",
	})

}
