package controllers

import (
	"backendtickitz/internal/models"
	"backendtickitz/internal/repositories"
	"backendtickitz/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileRepo *repositories.ProfileRepository
}

func NewProfileController(profileRepo *repositories.ProfileRepository) *ProfileController {
	return &ProfileController{profileRepo: profileRepo}
}

func (p *ProfileController) GetProfile(ctx *gin.Context) {

	payload, ok := ctx.Get("Payload")

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "silahkan login dahulu",
		})
		return
	}
	userClaims, ok := payload.(*pkg.Claims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "silahkan login dahulu",
		})
		return
	}
	log.Println("ini id userclaims", userClaims.Id)
	result, err := p.profileRepo.GetProfileById(ctx.Request.Context(), userClaims.Id)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan server",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success get profile by id",
		"data": result,
	})
}

func (p *ProfileController) UpdateProfileById(ctx *gin.Context) {
	newUser := &models.ProfileStruct{}
	if err := ctx.ShouldBind(newUser); err != nil {
		log.Println("sb", err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "terjadi kesalahan pada sistem",
		})
		return
	}
	payload, ok := ctx.Get("Payload")

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "silahkan login dahulu",
		})
		return
	}
	userClaims, ok := payload.(*pkg.Claims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "silahkan login dahulu",
		})
		return
	}

	result, err := p.profileRepo.UpdateProfile(ctx.Request.Context(), newUser, userClaims.Id)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan server",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success update profile by id",
		"data": result,
	})
}
