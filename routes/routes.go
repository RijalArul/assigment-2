package routes

import (
	"assigment-2/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)
	r.GET("/orders/:id", controllers.GetOrderByID)
	r.PUT("/orders/:orderId", controllers.UpdateOrder)
	r.DELETE("/orders/:orderId", controllers.DeleteOrder)
	r.Run()
}
