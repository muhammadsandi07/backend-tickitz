package controllers

import (
	"backendtickitz/internal/models"
	"backendtickitz/internal/repositories"
	"backendtickitz/pkg"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authRepo *repositories.AuthRepository
}

// initialization
func NewAuthController(authRepo *repositories.AuthRepository) *AuthController {
	return &AuthController{authRepo: authRepo}
}

// @Login
// @summary 			Login Users
// @router 				/auth [post]
// @accept				json
// @param				body body models.AuthStruct true "login information"
// @produce				json
// @version 			1.0
// @Failure 			400 {object} models.ErrorResponse "Bad Request"
// @Failure 			401 {object} models.ErrorResponse "Unauthorized"
// @Failure 			500 {object} models.ErrorResponse "Internal Server Error"
// @Success 			200 {object} models.AuthStruct "Success response"
func (a *AuthController) Register(ctx *gin.Context) {
	newUser := &models.AuthStruct{}
	if err := ctx.ShouldBind(newUser); err != nil {
		log.Println("error body", err)
		if status, msg := errorMsgBuilder(err); status != 0 {
			ctx.JSON(status, gin.H{
				"msg": msg,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "",
		})
		return
	}
	log.Println("[DEBUG BODY ]", newUser)
	var hash pkg.HashConfig
	hash.UseDefaultConfig()
	hashedPas, err := hash.GenHashedPassword(newUser.Password)
	if err != nil {

		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Hash failed",
		})
		return
	}

	result, err := a.authRepo.Register(ctx.Request.Context(), newUser, hashedPas)
	if err != nil {
		log.Println("error add user to db", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan sistem",
		})
		return
	}
	isEmailExist := result.ResultFirst.RowsAffected()

	if isEmailExist != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "email sudah ada",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "user berhasil ditambahkan",
		"data": newUser,
	})
}

func (a *AuthController) Login(ctx *gin.Context) {
	dataUser := &models.AuthStruct{}
	if err := ctx.ShouldBind(dataUser); err != nil {
		log.Println("error body", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan di sistem",
		})
		return
	}
	log.Println("[DEBUG EMAIL]", dataUser.Email)
	log.Println("[DEBUG PASSWORD]", dataUser.Password)
	var hash pkg.HashConfig
	result, err := a.authRepo.Login(ctx.Request.Context(), dataUser)
	log.Println("hassedh", result.Password)
	log.Println("ini result", result.Id, result.Password, result.Role)
	valid, err := hash.CompareHashAndPassword(result.Password, dataUser.Password)
	log.Println("valid", valid)
	if !valid {

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "silahkan login kembali",
		})
		return
	}
	claim := pkg.NewClaims(result.Id, result.Role)
	token, err := claim.GenerateToken()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "terjadi kesalahan server",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "Login success",
		"token": token,
	})

}

func errorMsgBuilder(err error) (status int, msg string) {
	if strings.Contains(err.Error(), "required") {
		return http.StatusBadRequest, "Email and password required"
	}
	return http.StatusBadRequest, "Email and password required"
}
