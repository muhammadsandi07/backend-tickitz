package controllers

import (
	"backendtickitz/internal/models"
	"backendtickitz/internal/repositories"
	"backendtickitz/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authRepo *repositories.AuthRepository
}

// initialization
func NewAuthController(authRepo *repositories.AuthRepository) *AuthController {
	return &AuthController{authRepo: authRepo}
}

// Register
// @summary					Register User
// @router					/auth/new [post]
// @accept					json
// @param					body body models.AuthSyruct true "register information"
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					201 {object} models.Response
func (a *AuthController) Register(ctx *gin.Context) {
	newUser := &models.AuthStruct{}
	if err := ctx.ShouldBind(newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}
	var hash pkg.HashConfig
	hash.UseDefaultConfig()
	hashedPas, err := hash.GenHashedPassword(newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}

	result, err := a.authRepo.Register(ctx.Request.Context(), newUser, hashedPas)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}
	isEmailExist := result.ResultFirst.RowsAffected()

	if isEmailExist != 0 {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Msg:  "Register Success",
		Data: result,
	})
}

// Login
// @summary					Login User
// @router					/auth [post]
// @accept					json
// @param					body body models.AuthStruct true "login information"
// @produce					json
// @failure					500 {object} models.ErrorResponse
// @success					200 {object} models.Response
func (a *AuthController) Login(ctx *gin.Context) {
	dataUser := &models.AuthStruct{}
	if err := ctx.ShouldBind(dataUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.InternalServerErrorCode,
				Status:  http.StatusInternalServerError,
				Message: "terjadi kesalahan server",
			},
		})
		return
	}

	var hash pkg.HashConfig
	result, err := a.authRepo.Login(ctx.Request.Context(), dataUser)
	valid, err := hash.CompareHashAndPassword(result.Password, dataUser.Password)
	if !valid {
		ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: &models.ErrorResponseDetail{
				Code:    models.UnAuthorized,
				Status:  http.StatusUnauthorized,
				Message: "terjadi kesalahan server",
			},
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
