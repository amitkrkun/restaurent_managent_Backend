package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-res/controllers"

	

)

func OrderItemRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.GET("/orderItem", controller.GetOrderItems())
	incomingRoutes.GET("/orderItem/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())	
	incomingRoutes.POST("/orderItem", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItem/:orderItem_id", controller.UpdateOrderItem() )

}