package routes

import (
	"backendtickitz/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Response struct {
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	// gin engine initialization
	router := gin.Default()
	movieRepo := repositories.NewMovieRepository(db)
	initPingRouter(router)
	initMovieRouter(router, movieRepo)
	return router
}
