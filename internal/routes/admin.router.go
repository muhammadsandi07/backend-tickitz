package routes

import (
	"backendtickitz/internal/controllers"
	"backendtickitz/internal/middlewares"
	"backendtickitz/internal/repositories"

	"github.com/gin-gonic/gin"
)

func initAdminRouter(router *gin.Engine, movieRepo *repositories.MovieRepository, middle *middlewares.Middlewares) {

	adminRouter := router.Group("/admin")
	{
		movieGroup := adminRouter.Group("/movies")
		adminController := controllers.NewAdminController(movieRepo)
		movieGroup.POST("", middle.VerifyToken, middle.AccessGate("admin"), adminController.AddMovie)
		movieGroup.PATCH("/:id", middle.VerifyToken, middle.AccessGate("admin"), adminController.UpdateMovie)
		movieGroup.DELETE("", middle.VerifyToken, middle.AccessGate("admin"), adminController.DeleteMovieById)

	}
}
