package routes

import (
	"github.com/gin-gonic/gin"
)

func initUsersRoute(router *gin.Engine) {
	// usersRouter := router.Group("/users")

	// login
	// usersRouter.POST("/auth", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"msg":  "user ditemukan",
	// 		"data": users,
	// 	})
	// })
	// register
	// usersRouter.POST("/auth/new", func(ctx *gin.Context) {
	// 	newUser := &models.UserStruct{}
	// 	if err := ctx.ShouldBind(newUser); err != nil {
	// 		log.Println(err)
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{
	// 			"msg": "terjadi kesalahan sistem",
	// 		})
	// 		return
	// 	}

	// 	for _, user := range users {
	// 		isEmailExist := user.Email == newUser.Email
	// 		if isEmailExist {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{
	// 				"msg": "User sudah tersedia",
	// 			})
	// 			return
	// 		}
	// 		newUsers := append(users, *newUser)
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"msg":  "account berhasil di buat",
	// 			"data": newUsers,
	// 		})
	// 	}
	// })

}
