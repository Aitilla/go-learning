package db

import (
    "context"
    "log"
    "os"

    "github.com/jackc/pgx/v5"
    "github.com/joho/godotenv"
)

var Conn *pgx.Conn

func Init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    url := os.Getenv("DATABASE_URL") 
    if url == "" {
        log.Fatal("INVALID DB URL")
    }

    Conn, err = pgx.Connect(context.Background(), url)
    if err != nil {
        log.Fatal("Failed to connect:", err)
    }
}
