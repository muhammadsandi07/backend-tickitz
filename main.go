package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	type userStruct struct {
		Id       int    `json:"id" form:"id"`
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
	type movie struct {
		Id   int
		Name string
	}

	users := []userStruct{
		{Id: 1, Email: "sandi@gmail.com", Password: "12345678"},
		{Id: 2, Email: "denis@gmail.com", Password: "12345678"},
		{Id: 3, Email: "rohman@gmail.com", Password: "12345678"},
		{Id: 4, Email: "sandi@gmail.com", Password: "12345678"},
	}

	// login
	router.POST("/auth", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "user ditemukan",
			"data": users,
		})
	})
	// register
	router.POST("/auth/new", func(ctx *gin.Context) {
		newUser := &userStruct{}
		if err := ctx.ShouldBind(newUser); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "terjadi kesalahan sistem",
			})
			return
		}

		for _, user := range users {
			accountAlready := strings.EqualFold(user.Email, newUser.Email)

			if accountAlready {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "User sudah tersedia",
				})
				return
			}
			newUsers := append(users, *newUser)
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "account berhasil di buat",
				"data": newUsers,
			})
		}
	})
	// movie

	type movieStruct struct {
		Id       int    `json:"id" from:"id"`
		Name     string `json:"name" from:"name"`
		duration int    `json:"duration" from:"duration"`
		synopsis string `json:"synopsis" from:"synopsis"`
	}

	movies := []movieStruct{
		{Id: 1, Name: "spiderman", duration: 200, synopsis: "ini spiderman"},
		{Id: 2, Name: "suster kagak ngesot", duration: 200, synopsis: "ini mah kayang kayanya"},
		{Id: 3, Name: "pocong kepeleset", duration: 200, synopsis: "kagak bisa diri lagi"},
		{Id: 4, Name: "tuyul sedekah", duration: 200, synopsis: "tuyul versi ramadhan"},
		{Id: 5, Name: "jelangkung minta anter", duration: 200, synopsis: "jelangkung minta anter"},
	}

	router.GET("/movies", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "get all movie success",
			"data": movies,
		})
	})
	router.GET("/movies/:id", func(ctx *gin.Context) {
		idMovie, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "params is is needed",
			})
		}

		idInt, err := strconv.Atoi(idMovie)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "terjadi kesalahan server",
			})
			return
		}

		var movie []movieStruct
		for _, m := range movies {
			if m.Id == idInt {
				movie = append(movie, m)
				break
			}
		}
		if len(movie) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "movie tidak tersedia",
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "success get movie by id",
			"data": movie,
		})

	})

	router.Run("127.0.0.1:8080")
}
