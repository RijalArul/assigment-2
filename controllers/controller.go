package controllers

import (
	"assigment-2/config"
	"assigment-2/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	DB := config.GetDB()
	var errTx error
	tx := DB.Begin()
	currentTime := time.Now()

	customer_name, item_code, desc, quantity := ctx.PostForm("customer_name"), ctx.PostForm("item_code"), ctx.PostForm("desc"), ctx.PostForm("quantity")
	quantityParse, _ := strconv.ParseUint(quantity, 10, 32)
	payloadOrder := &models.Order{
		CustomerName: customer_name,
		OrderedAt:    currentTime,
	}

	if err := tx.Create(&payloadOrder).Error; err != nil {
		errTx = err

	}

	payloadItem := &models.Item{
		ItemCode:    item_code,
		Description: desc,
		Quantity:    int(quantityParse),
		OrderID:     int(payloadOrder.ID),
	}

	if err := tx.Create(&payloadItem).Error; err != nil {
		errTx = err
		// tx.Rollback()
	}

	if errTx != nil {
		ctx.JSON(400, gin.H{
			"message": "Failed Create",
			"err":     errTx,
		})
		tx.Rollback()
	} else {
		var RespCodeSuccess []map[string]interface{}
		RespCodeSuccess = []map[string]interface{}{
			{
				"itemCode": payloadItem.ItemCode,
				"desc":     payloadItem.Description,
				"quantity": payloadItem.Quantity,
			},
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"orderedAt":    payloadOrder.OrderedAt,
			"customerName": payloadOrder.CustomerName,
			"items":        RespCodeSuccess,
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
	orderId := ctx.Param("orderId")
	customer_name, item_code, desc, quantity := ctx.PostForm("customer_name"), ctx.PostForm("item_code"), ctx.PostForm("desc"), ctx.PostForm("quantity")
	quantityParse, _ := strconv.ParseUint(quantity, 10, 32)
	currentTime := time.Now()
	var errMsg error
	order := models.Order{}
	item := models.Item{}

	if err := config.GetDB().Model(&order).Where("id = ?", orderId).Updates(&models.Order{CustomerName: customer_name, OrderedAt: currentTime}).First(&order).Error; err != nil {
		errMsg = err
	}
	config.GetDB().Model(&item).Where("items.order_id", orderId).Updates(&models.Item{ItemCode: item_code, Description: desc, Quantity: int(quantityParse)}).First(&item)

	if errMsg != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": errMsg,
		})
	} else {
		itemsUpdate := []map[string]interface{}{
			{
				"lineItemId":  item.ID,
				"itemCode":    item.ItemCode,
				"description": item.Description,
				"quantity":    item.Quantity,
			},
		}

		ctx.JSON(http.StatusOK, gin.H{
			"orderId":      order.ID,
			"customerName": order.CustomerName,
			"orderedAt":    order.OrderedAt,
			"items":        itemsUpdate,
		})
	}
}

func DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	order := models.Order{}

	config.GetDB().Delete(&order, orderId)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success deleted order",
	})

}
