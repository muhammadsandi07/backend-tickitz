package controllers

import (
	"backendtickitz/internal/repositories"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	scheduleRepo *repositories.ScheduleRepository
}

func NewScheduleController(scheduleRepo *repositories.ScheduleRepository) *ScheduleController {
	return &ScheduleController{scheduleRepo: scheduleRepo}
}

func (s *ScheduleController) GetSchedule(ctx *gin.Context) {
	idMovie, ok := ctx.Params.Get("id")
	location := ctx.Query("location")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "params id is needed",
		})
		return
	}
	log.Println(location, "ini location dari query")
	log.Println("ini id movie", idMovie)
	idMovie = strings.TrimPrefix(idMovie, ":")
	idInt, err := strconv.Atoi(idMovie)
	result, err := s.scheduleRepo.GetSchedule(ctx.Request.Context(), idInt, location)
	if err != nil {
		log.Println("[debug schedule get]", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "terjadi kesalahan server",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success get schedule by id movie",
		"data": result,
	})
}
