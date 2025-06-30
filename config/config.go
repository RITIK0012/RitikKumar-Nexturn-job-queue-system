package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func ConnectDB() *sql.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Ping failed: %v", err)
	}
	return db
}

func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}
