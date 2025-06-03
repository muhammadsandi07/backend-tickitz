package controllers

import (
	"backendtickitz/internal/models"
	"backendtickitz/internal/repositories"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	movieRepo *repositories.MovieRepository
}

// initialization
func NewMovieController(movieRepo *repositories.MovieRepository) *MovieController {
	return &MovieController{movieRepo: movieRepo}
}

func (m *MovieController) GetMovies(ctx *gin.Context) {
	var params models.MovieQueryParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	log.Println("ini data params", params)

	result, err := m.movieRepo.GetMovies(ctx.Request.Context(), &params)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan sistem",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": result,
	})
	return

}

func (m *MovieController) GetMovieById(ctx *gin.Context) {
	idMovie, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "params id is needed",
		})
		return
	}
	idMovie = strings.TrimPrefix(idMovie, ":")
	idInt, err := strconv.Atoi(idMovie)
	log.Println("[LOG PRINT ID MOVIE 1]", idMovie)
	log.Println("[LOG PRINT ID MOVIE 2]", idInt)
	result, err := m.movieRepo.GetMovieById(ctx.Request.Context(), idInt)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan server",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success get movie by id",
		"data": result,
	})
}

func (m *MovieController) GetMovieUpcoming(ctx *gin.Context) {
	result, err := m.movieRepo.Upcoming(ctx.Request.Context())
	if err != nil {
		log.Println("[DEBUG] ini error upcoming", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan server",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success get movie by id",
		"data": result,
	})
}
func (m *MovieController) GetMoviePopular(ctx *gin.Context) {
	result, err := m.movieRepo.Popular(ctx.Request.Context())
	if err != nil {
		log.Println("[DEBUG] ini error popular", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan server",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success get movie ",
		"data": result,
	})
}
