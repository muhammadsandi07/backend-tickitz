package controllers

import (
	"backendtickitz/internal/models"
	"backendtickitz/internal/repositories"
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

// Get Movies
// @summary					Get All Movies
// @router					/movies [GET]
// @accept					json
// @param					query query  models.MovieQueryParams false "Query Parameters"
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					200 {object} models.Response
func (m *MovieController) GetMovies(ctx *gin.Context) {
	var params models.MovieQueryParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.BadRequest,
				Status:  http.StatusBadRequest,
				Message: "Inputan User failed",
			},
		})
		return
	}

	result, err := m.movieRepo.GetMovies(ctx.Request.Context(), &params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "get data movies failed",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": result,
	})
	return

}

// Get Movies By ID
// @summary					Get Movies By ID
// @router					/movies/:id [GET]
// @accept					json
// @param					id path  int true "Query Parameters"
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					200 {object} models.Response
func (m *MovieController) GetMovieById(ctx *gin.Context) {
	idMovie, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.BadRequest,
				Status:  http.StatusBadRequest,
				Message: "Params id needed",
			},
		})
		return
	}
	idMovie = strings.TrimPrefix(idMovie, ":")
	idInt, err := strconv.Atoi(idMovie)
	result, err := m.movieRepo.GetMovieById(ctx.Request.Context(), idInt)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "Get Movie By Id Failed",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Msg:  "Get Movie By ID Success",
		Data: result,
	})
}

// Get Movies Up Coming
// @summary					Get Movies Up Coming
// @router					/movies/upcoming [GET]
// @accept					json
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					200 {object} models.Response
func (m *MovieController) GetMovieUpcoming(ctx *gin.Context) {
	result, err := m.movieRepo.Upcoming(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "Get Movie UpComing Failed",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Msg:  "Get Movie Up Coming Success",
		Data: result,
	})
}

// Get Movies Popular
// @summary					Get Movies Popular
// @router					/movies/popular [GET]
// @accept					json
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					200 {object} models.Response
func (m *MovieController) GetMoviePopular(ctx *gin.Context) {
	result, err := m.movieRepo.Popular(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "Get Movie Popular Failed",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Msg:  "Get Movie Up Coming",
		Data: result,
	})
}
