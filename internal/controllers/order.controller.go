package controllers

import (
	"backendtickitz/internal/models"
	"backendtickitz/internal/repositories"
	"backendtickitz/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderRepo *repositories.OrderRepository
}

func NewOrderController(orderRepo *repositories.OrderRepository) *OrderController {
	return &OrderController{orderRepo: orderRepo}
}

func (o *OrderController) AddOrder(ctx *gin.Context) {

	newOrder := &models.OrderStruct{}
	if err := ctx.ShouldBind(&newOrder); err != nil {
		log.Println("[DEBUG ERROR]", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "terjadi kesalahan pada sistem",
		})
		return
	}

	claims, ok := ctx.Get("Payload")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "silahkan login dahulu",
		})
		return
	}
	usersClaims, ok := claims.(*pkg.Claims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{

			"message": "Identitas login anda rusak, Silahkan login kembali",
		})
		return
	}
	newOrder.IdUser = usersClaims.Id
	log.Println("[DEBUG new order]", newOrder)
	log.Println("[DEBUG id user]", usersClaims)
	err := o.orderRepo.AddOrder(ctx.Request.Context(), newOrder)
	log.Println("err", err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "gagal add order",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "add order success",
	})

}

func (o *OrderController) GetOrderByUser(ctx *gin.Context) {

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
	result, err := o.orderRepo.GetOrderByUser(ctx.Request.Context(), userClaims.Id)
	if err != nil {
		log.Println("[DEBUG]", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan server",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success get order by id",
		"data": result,
	})
}
