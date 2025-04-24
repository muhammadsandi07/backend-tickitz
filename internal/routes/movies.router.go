package routes

import (
	"backendtickitz/internal/controllers"
	"backendtickitz/internal/models"
	"backendtickitz/internal/repositories"
	"backendtickitz/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initMovieRouter(router *gin.Engine, movieRepo *repositories.MovieRepository) {
	movieRouter := router.Group("/movies")
	movieController := controllers.NewMovieController(movieRepo)
	movieRouter.GET("", movieController.GetMovies)
	movieRouter.GET("/:id", movieController.GetMovieById)

	movieRouter.POST("/movies", func(ctx *gin.Context) {
		newMovie := &models.MovieStruct{}
		if err := ctx.ShouldBind(newMovie); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "terjadi kesalahan sistem",
			})
			return
		}

		query := "INSERT INTO movie (name, duration, synopsis, img_movie, backdrop, release_date) VALUES ($1, $2, $3, $4, $5, $6)"
		value := []any{newMovie.Name, newMovie.Duration, newMovie.Synopsis, newMovie.Img_movie, newMovie.Backdrop, newMovie.Release_Date}

		cmd, err := pkg.DB.Exec(ctx.Request.Context(), query, value...)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "terjadi kesalahan server",
			})
			return
		}
		if cmd.RowsAffected() == 0 {
			log.Println("query gagal, tidak merubah data di DB")
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
}
