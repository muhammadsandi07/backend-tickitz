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
		cinemaGroup := adminRouter.Group("/cinema")
		movieGroup := adminRouter.Group("/movies")
		adminController := controllers.NewAdminController(movieRepo)
		movieGroup.GET("/genres", middle.VerifyToken, middle.AccessGate("admin"), adminController.GetGenre)
		movieGroup.POST("", middle.VerifyToken, middle.AccessGate("admin"), adminController.AddMovie)
		movieGroup.PATCH("/:id", middle.VerifyToken, middle.AccessGate("admin"), adminController.UpdateMovie)
		movieGroup.DELETE("", middle.VerifyToken, middle.AccessGate("admin"), adminController.DeleteMovieById)
		cinemaGroup.GET("", middle.VerifyToken, middle.AccessGate("admin"), adminController.GetCinema)
	}
}
