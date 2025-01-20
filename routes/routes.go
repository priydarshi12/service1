package routes

import (
	"github.com/gin-gonic/gin"
	"myapp/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/order/create", controllers.CreateOrder)
	router.POST("/order/create/bulk", controllers.BulkOrder)
	router.GET("/order/get", controllers.GetOrder)
	router.PUT("/order/update", controllers.UpdateOrder)
	router.DELETE("/order/delete", controllers.DeleteOrder)
}
