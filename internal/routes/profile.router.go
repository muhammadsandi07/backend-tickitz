package routes

import (
	"backendtickitz/internal/controllers"
	"backendtickitz/internal/middlewares"
	"backendtickitz/internal/repositories"

	"github.com/gin-gonic/gin"
)

func initProfileRouter(router *gin.Engine, profileRepo *repositories.ProfileRepository, middle *middlewares.Middlewares) {
	profileRouter := router.Group("/profile")
	profileController := controllers.NewProfileController(profileRepo)
	profileRouter.GET("", middle.VerifyToken, profileController.GetProfile)
	profileRouter.PATCH("", middle.VerifyToken, profileController.UpdateProfileById)
}
