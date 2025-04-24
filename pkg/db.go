package pkg

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() (*pgxpool.Pool, error) {
	// create database connection string
	dbEnv := []any{}
	dbEnv = append(dbEnv, os.Getenv("DBUSER"))
	dbEnv = append(dbEnv, os.Getenv("DBPASS"))
	dbEnv = append(dbEnv, os.Getenv("DBHOST"))
	dbEnv = append(dbEnv, os.Getenv("DBPORT"))
	dbEnv = append(dbEnv, os.Getenv("DBNAME"))
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv...)
	var err error
	DB, err = pgxpool.New(context.Background(), dbString)
	if err != nil {
		return nil, err
	}
	// test ping db

	err = DB.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Ping failed: %w", err)
	}
	log.Println("Connected to PostgreSQL")
	return DB, nil
}
