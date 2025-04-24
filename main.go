package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	router := gin.Default()
	dbEnv := []any{}
	dbEnv = append(dbEnv, os.Getenv("DBUSER"))
	dbEnv = append(dbEnv, os.Getenv("DBPASS"))
	dbEnv = append(dbEnv, os.Getenv("DBHOST"))
	dbEnv = append(dbEnv, os.Getenv("DBPORT"))
	dbEnv = append(dbEnv, os.Getenv("DBNAME"))
	// setup database connection
	for _, v := range dbEnv {
		log.Println("[debug]", v)
	}
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv...)
	dbClient, err := pgxpool.New(context.Background(), dbString)
	if err != nil {
		log.Printf("Unable to create connection pool:%v\n", err)
		os.Exit(1)
	}
	defer func() {
		log.Println("Closing DB...")
		dbClient.Close()
	}()
	// endpoint & resource
	// /ping => protocol://hostname/ping => http://localhost:port/ping
	router.GET("/ping", func(ctx *gin.Context) {

		type students struct {
			Id   int    "json: id"
			name string "json: id"
		}
		// query := "SELECT id, name from students"
		// rows, err := dbClient.Query(context.Background(), query)
		// if err != nil{
		// 	log.Println(err.Error())
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{
		// 		"msg": "terjadi kesalahan sistem",
		// 	})
		// 	return
		// }

		// defer rows.Close()
		// var result

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	type userStruct struct {
		Id       int    `json:"id" form:"id"`
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
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
			isEmailExist := user.Email == newUser.Email
			if isEmailExist {
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
		Id           int       `json:"id,omitempty" form:"id" db:"id"`
		Name         string    `json:"name" form:"name" db:"name"`
		Duration     int       `json:"duration" form:"duration" db:"duration"`
		Synopsis     string    `json:"synopsis" form:"synopsis" db:"synopsis"`
		Img_movie    string    `json:"img_movie" form:img_movie db:"img_movie"`
		Backdrop     string    `json:"backdrop" form:backdrop db:"backdrop"`
		Release_Date time.Time `json:"release_date" form:release_date db:"release_date"`
	}

	router.GET("/movies", func(ctx *gin.Context) {
		nameQ := ctx.Query("name")
		// type Movie struct {
		// 	Id   int    `json:"id" db:"id"`
		// 	Name string `json:"name" db:"name"`
		// }

		if nameQ == "" {
			query := "select id, name, duration, synopsis, img_movie, backdrop, release_date FROM movie"
			rows, err := dbClient.Query(ctx.Request.Context(), query)
			if err != nil {
				log.Println(err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "terjadi kesalahan sistem",
				})
				return
			}
			defer rows.Close()
			var result []movieStruct
			for rows.Next() {
				var movie movieStruct
				if err := rows.Scan(&movie.Id, &movie.Name, &movie.Duration, &movie.Synopsis, &movie.Img_movie, &movie.Backdrop, &movie.Release_Date); err != nil {
					log.Println(err.Error())
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"msg": "terjadi kesalahan sistem",
					})
					return
				}
				result = append(result, movie)
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
	})
	router.GET("/movies/:id", func(ctx *gin.Context) {

		idMovie, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "params id is needed",
			})
		}

		idInt, err := strconv.Atoi(idMovie)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "terjadi kesalahan server",
			})
			return
		}
		query := "select id, name, duration, synopsis, img_movie, backdrop, release_date FROM movie where id =$1"
		var result movieStruct
		if err := dbClient.QueryRow(ctx.Request.Context(), query, idInt).Scan(&result.Id, &result.Name, &result.Duration, &result.Synopsis, &result.Img_movie, &result.Backdrop, &result.Release_Date); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "terjadi kesalahan server",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "success get movie by id",
			"data": result,
		})
	})

	router.POST("/movies", func(ctx *gin.Context) {
		newMovie := &movieStruct{}
		if err := ctx.ShouldBind(newMovie); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "terjadi kesalahan sistem",
			})
			return
		}

		query := "INSERT INTO movie (name, duration, synopsis, img_movie, backdrop, release_date) VALUES ($1, $2, $3, $4, $5, $6)"
		value := []any{newMovie.Name, newMovie.Duration, newMovie.Synopsis, newMovie.Img_movie, newMovie.Backdrop, newMovie.Release_Date}

		cmd, err := dbClient.Exec(ctx.Request.Context(), query, value...)
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
	router.Run("127.0.0.1:8085")
}
