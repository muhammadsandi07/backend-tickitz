package routes

import (
	"backendtickitz/internal/controllers"
	"backendtickitz/internal/middlewares"
	"backendtickitz/internal/repositories"

	"github.com/gin-gonic/gin"
)

func initScheduleRouter(router *gin.Engine, scheduleRepo *repositories.ScheduleRepository, middle *middlewares.Middlewares) {
	scheduleRouter := router.Group("/schedule")
	scheduleController := controllers.NewScheduleController(scheduleRepo)
	scheduleRouter.GET("/:id", scheduleController.GetSchedule)
}
