package db_postgresql

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// Config holds DB config - you can extend this or read from env/JSON/config
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

// NewConfigFromEnv returns DB config from environment variables
func NewConfigFromEnv() Config {
	return Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     5432, // default port
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable", // for local dev; use "require" in prod
		TimeZone: "UTC",
	}
}

// Connect initializes and returns a singleton *gorm.DB
func Connect(cfg Config) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode, cfg.TimeZone,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // optional: for query logs
		})
		if err != nil {
			log.Fatalf("failed to connect to PostgreSQL: %v", err)
		}

		// Optional: configure connection pool
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("failed to get DB from GORM: %v", err)
		}
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(25)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)

		dbInstance = db
	})

	return dbInstance
}
