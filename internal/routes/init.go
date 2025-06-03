package routes

import (
	"backendtickitz/internal/middlewares"
	"backendtickitz/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Response struct {
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

func InitRouter(db *pgxpool.Pool, rdb *redis.Client) *gin.Engine {
	// gin engine initialization
	router := gin.Default()
	movieRepo := repositories.NewMovieRepository(db, rdb)
	authRepo := repositories.NewAuthRepository(db)
	profileRepo := repositories.NewProfileRepostory(db)
	scheduleRepo := repositories.NewScheduleRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	middlewares := middlewares.InitMiddleware()
	router.Use(middlewares.CORSMiddleware)
	router.Static("img_movie", "./public/img_movie")
	initPingRouter(router)
	initAdminRouter(router, movieRepo, middlewares)
	initAuthRoute(router, authRepo)
	initMovieRouter(router, movieRepo, middlewares)
	initProfileRouter(router, profileRepo, middlewares)
	initScheduleRouter(router, scheduleRepo, middlewares)
	initOrderRouter(router, orderRepo, middlewares)
	return router
}
