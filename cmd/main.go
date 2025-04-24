package main

import (
	"backendtickitz/internal/routes"
	"backendtickitz/pkg"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

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

	router := routes.InitRouter(pg)
	router.Run("127.0.0.1:8085")
}
