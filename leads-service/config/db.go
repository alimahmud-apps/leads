package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is a global database connection pool.
var DB *sqlx.DB

// InitDB initializes the database connection.
func InitDB() {
	var err error
	// err = godotenv.Load(
	// 	"../../.env", // First try the project root directory
	// 	".env",       // Then try the current working directory inside the container
	// )
	// if err != nil {
	// 	log.Println("No .env file found, using default config")
	// }
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("DB_HOST_SERVICE", "localhost"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_USER", "wpuser"),
		GetEnv("DB_PASSWORD", "wppassword"),
		GetEnv("DB_NAME", "wordpress"),
	)

	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	log.Println("Successfully connected to the database!")
}

// GetEnv retrieves environment variables or returns a default value.
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	fmt.Println("value : ", value)
	if !exists {
		return defaultValue
	}
	return value
}
