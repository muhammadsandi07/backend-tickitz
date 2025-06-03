package routes

import (
	"backendtickitz/internal/controllers"
	"backendtickitz/internal/middlewares"
	"backendtickitz/internal/repositories"

	"github.com/gin-gonic/gin"
)

func initMovieRouter(router *gin.Engine, movieRepo *repositories.MovieRepository, middle *middlewares.Middlewares) {
	movieRouter := router.Group("/movies")
	movieController := controllers.NewMovieController(movieRepo)
	movieRouter.GET("", movieController.GetMovies)
	movieRouter.GET("/:id", movieController.GetMovieById)
	movieRouter.GET("/upcoming", movieController.GetMovieUpcoming)
	movieRouter.GET("/popular", movieController.GetMoviePopular)
}
