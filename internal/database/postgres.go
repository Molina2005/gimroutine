package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error al cargar archivo .env", err)
	}
	connectionChain := fmt.Sprintf(
		"user=%v password=%v dbname=%v host=%v port=%v sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	pool, err := pgxpool.New(context.Background(), connectionChain)
	if err != nil {
		log.Fatal("error conectando a la base de datos", err)
	}
	return pool, nil
}
