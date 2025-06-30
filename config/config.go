package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sqlx.DB

func ConnectDB() *sqlx.DB {
	dsn := os.Getenv("DATABASE_URL") // use Render’s env variable
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Postgres connection failed: %v", err)
	}
	DB = db
	return DB
}

func InitSchema() {
	query := `
	CREATE TABLE IF NOT EXISTS jobs (
		id SERIAL PRIMARY KEY,
		payload TEXT,
		status VARCHAR(50),
		result TEXT,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create jobs table: %v", err)
	}
}

func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}
