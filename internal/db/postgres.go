package db

import (
	"context"
	"fmt"
	"log"

	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	var dbError error
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, database)
	DB, dbError = pgxpool.New(context.Background(), connString)
	if dbError != nil {
		log.Fatal("Error connecting to the database.", err)
	}
	fmt.Println("Connected to the database.")
}

func Close() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed.")
	}
}
