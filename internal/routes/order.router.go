package routes

import (
	"backendtickitz/internal/controllers"
	"backendtickitz/internal/middlewares"
	"backendtickitz/internal/repositories"

	"github.com/gin-gonic/gin"
)

func initOrderRouter(router *gin.Engine, orderRepo *repositories.OrderRepository, middleware *middlewares.Middlewares) {
	orderRouter := router.Group("/order")
	orderController := controllers.NewOrderController(orderRepo)
	orderRouter.POST("", middleware.VerifyToken, orderController.AddOrder)
	orderRouter.GET("", middleware.VerifyToken, orderController.GetOrderByUser)
}
