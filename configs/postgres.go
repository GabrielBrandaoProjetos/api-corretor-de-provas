package configs

import (
	"context"
	"fmt"
	"log"

	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/repository"
	"github.com/GabrielBrandaoProjetos/corretor-de-provas-api/internal/util/stringutil"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func InitPostgres() *repository.Queries {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		stringutil.GetEnv("DB_HOST", "db"),
		stringutil.GetEnv("DB_PORT", "5432"),
		stringutil.GetEnv("DB_USER", "postgres"),
		stringutil.GetEnv("DB_PASSWORD", "postgres"),
		stringutil.GetEnv("DB_NAME", "corrector"),
		stringutil.GetEnv("DB_SSL_MODE", "disable"),
	)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil
	}
	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal("Error ping to database: ", err)
		return nil
	}
	repository := repository.New(conn)

	return repository
}
