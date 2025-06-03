package routes

import (
	"backendtickitz/internal/controllers"
	"backendtickitz/internal/repositories"

	"github.com/gin-gonic/gin"
)

func initAuthRoute(router *gin.Engine, authRepo *repositories.AuthRepository) {
	authRouter := router.Group("/auth")
	authController := controllers.NewAuthController(authRepo)
	authRouter.POST("", authController.Login)
	authRouter.POST("/new", authController.Register)

}
