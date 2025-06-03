package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]

func initPingRouter(router *gin.Engine) {
	userRouter := router.Group("/ping")
	userRouter.GET("", func(ctx *gin.Context) {

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

}
