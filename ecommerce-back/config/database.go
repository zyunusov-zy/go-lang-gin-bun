package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)


var DB *bun.DB

func LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No Port in env")
		return ""
	}
	return port
}

func ConnectDB() {
	LoadEnv()
	dsn :=os.Getenv("DB_URL")
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}
	DB = bun.NewDB(sqlDB, pgdialect.New())
	log.Println("Connected to PostgreSQL using Bun ORM")
}