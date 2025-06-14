package controllers

import (
	"backendtickitz/internal/models"
	"backendtickitz/internal/repositories"
	"backendtickitz/pkg"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	fp "path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	movieRepo *repositories.MovieRepository
}

func NewAdminController(movieRepo *repositories.MovieRepository) *AdminController {
	return &AdminController{movieRepo: movieRepo}
}

// Add Movie
// @summary					Add Movie
// @router					/movies [POST]
// @accept					json
// @param					id path  int true "Query Parameters"
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					200 {object} models.Response
func (a *AdminController) AddMovie(ctx *gin.Context) {
	// newMovie := &models.MovieStruct{}
	// file, err1 := ctx.FormFile("img")
	// if err1 != nil {
	// 	log.Println(err1.Error())
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Terjadi kesalahan server",
	// 	})
	// }
	var formBody models.MovieFrom
	if err := ctx.ShouldBind(&formBody); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "terjadi kesalahan sistem",
		})
		return
	}
	uploadedFiles := map[string]any{
		"img_movie": "",
		"backdrop":  "",
	}

	uploadFile := func(file *multipart.FileHeader, key string) {
		if file != nil {
			filename, _, err := fileHandling(ctx, file)
			if err != nil {
				log.Printf("Gagal upload %s: %v,", key, err)
				return
			}
			uploadedFiles[key] = filename
		}
	}
	uploadFile(formBody.Img_movie, "img_movie")
	uploadFile(formBody.Backdrop, "backdrop")

	err := a.movieRepo.AddMovie(ctx.Request.Context(), &formBody, uploadedFiles)
	log.Println("[DEBUG ERROR]", err)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan sistem",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Add Movie Success",
	})

}

func fileHandling(ctx *gin.Context, file *multipart.FileHeader) (filename, filepath string, err error) {
	claims, _ := ctx.Get("Payload")
	userClaims := claims.(*pkg.Claims)
	log.Println("ini payload", userClaims)
	ext := fp.Ext(file.Filename)
	filename = fmt.Sprintf("%d_%d_movie_image%s", time.Now().UnixNano(), userClaims.Id, ext)
	filepath = fp.Join("public", "img_movie", filename)

	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Terjadi kesalahan upload",
		})
		return "", "", nil
	}
	return filename, filepath, nil
}

// Update Movie
// @summary					Update Movie
// @router					/movies/:id [PATCH]
// @accept					json
// @param					id path  int true "Query Parameters"
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					200 {object} models.Response
func (a *AdminController) UpdateMovie(ctx *gin.Context) {
	newMovie := &models.MovieStruct{}
	idMovie, ok := ctx.Params.Get("id")
	if err := ctx.ShouldBind(newMovie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "terjadi kesalahan pada sistem",
		})
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "param id is needed",
		})
		return
	}
	idInt, _ := strconv.Atoi(idMovie)
	err := a.movieRepo.UpdateMovie(ctx.Request.Context(), newMovie, idInt)
	if err != nil {
		log.Printf("[debug err]: %+v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "update movie berhasil",
	})
	return
}

// Delete Movie
// @summary					Delete Movie
// @router					/movies/:id [DELETE]
// @accept					json
// @param					id path  int true "Query Parameters"
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					200 {object} models.Response
func (a *AdminController) DeleteMovieById(ctx *gin.Context) {
	var movies models.MovieStruct
	err := ctx.ShouldBind(&movies)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body request",
		})
		return
	}
	log.Println("[DEBUG MOVIES]", movies)
	result, err := a.movieRepo.DeleteMovie(ctx.Request.Context(), &movies)
	if err != nil {
		log.Println("anak ayam", err)
	}
	row := result.ResultSecond.RowsAffected()
	// log.Println("haha hihi", row)

	if row == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "movie tidak tersedia",
		})
		return
	}
	if err != nil {
		log.Println("error1", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalah di server",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "deleted movie success",
	})

}

func (a *AdminController) GetGenre(ctx *gin.Context) {
	result, err := a.movieRepo.GetGenres(ctx.Request.Context())
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
func (a *AdminController) GetCinema(ctx *gin.Context) {
	result, err := a.movieRepo.GetCinema(ctx.Request.Context())
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
