package main

import (
	_ "backendtickitz/docs"
	"backendtickitz/internal/routes"
	"backendtickitz/pkg"
	"log"
	"os"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/joho/godotenv/autoload"
)

// @title 				TICKITZ API
// @version 			1.0
// @description 		example of working Backedn crated during class
// @host				localhose:8080
// @BasePath			/

func main() {

	// setup database connection
	pg, err := pkg.Connect()
	if err != nil {
		log.Printf("[ERROR] unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		log.Println("Closing DB...")
		pg.Close()
	}()
	// endpoint & resource
	// /ping => protocol://hostname/ping => http://localhost:port/ping

	// movie
	password := "fazztrack"
	var has pkg.HashConfig
	has.UseDefaultConfig()
	result, _ := has.GenHashedPassword(password)
	log.Println("password", password)
	log.Println("passwordHashed", result)
	rdb := pkg.RedisConnect()
	router := routes.InitRouter(pg, rdb)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("127.0.0.1:8080")

}
