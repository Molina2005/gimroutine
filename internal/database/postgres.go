package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {
	// cargue de archivo .env
	_ = godotenv.Load()

	connectionChain := fmt.Sprintf(
		// Formato que pide supabase para la conexion
		"postgres://%s:%s@%s:%s/%s?sslmode=require",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	// Se desactiva configuracion de protocolo avanzado que trae pgxpool para las consultas
	// Y se pasa al protocola simple el cual solo envia la consulta y ya
	config, err := pgxpool.ParseConfig(connectionChain)
	if err != nil {
		log.Fatal("error al obtener config", err)
	}
	// Se pasa al nuevo protocola simple
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error conectando a la base de datos", err)
	}
	return pool, nil
}
