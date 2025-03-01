package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	// Load environment variables
	err := godotenv.Load(
		"../../.env", // First try the project root directory
		".env",       // Then try the current working directory inside the container
	)
	if err != nil {
		log.Println("No .env file found, using default config")
	}

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("DB_HOST_SERVICE", "localhost"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_USER", "wpuser"),
		GetEnv("DB_PASSWORD", "wppassword"),
		GetEnv("DB_NAME", "wordpress"),
	)

	// Connect to the database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := goose.Up(db, "sql"); err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Migration completed successfully!")
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
