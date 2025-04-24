package controllers

import (
	"backendtickitz/internal/repositories"
	"log"
	"net/http"
	"strconv"

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
	nameQ := ctx.Query("name")
	if nameQ == "" {
		result, err := m.movieRepo.GetMovies(ctx.Request.Context())
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "terjadi kesalahan sistem",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": result,
		})
		return
	}
	// query := "SELECT id, name from movie where name like $1"
	// values := []any{nameQ}
	// var result Movie
	// if err := dbClient.QueryRow(context.Background(), query, values...).Scan(&result.Id, &result.Name); err != nil {
	// 	log.Println(err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"msg": "Terjadi kesalahan sistem",
	// 	})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"msg":  "Success",
	// 	"user": result,
	// })
	// for _, m := range movies {
	// 	condition := strings.EqualFold(m.Name, nameQ)
	// 	if condition {
	// 		result = append(result, m)
	// 	}

	// }
	// if len(result) == 0 {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"msg": "Movie tidak ditemukan ",
	// 	})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"msg":  "get movie success",
	// 	"data": result,
	// })
}

func (m *MovieController) GetMovieById(ctx *gin.Context) {
	idMovie, ok := ctx.Params.Get("id")

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "params id is needed",
		})
	}
	idInt, err := strconv.Atoi(idMovie)
	result, err := m.movieRepo.GetMovieById(ctx.Request.Context(), idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan server",
			"err": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success get movie by id",
		"data": result,
	})
}
