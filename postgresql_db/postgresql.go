package postgresql_db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	dsn := "host=localhost user=root password=11111111 dbname=cyberus_db port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// Get generic database object sql.DB to configure pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	// ✅ Configure connection pool
	sqlDB.SetMaxOpenConns(25)                  // max open connections
	sqlDB.SetMaxIdleConns(10)                  // max idle connections
	sqlDB.SetConnMaxLifetime(30 * time.Second) // max lifetime of a connection

	// Test connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("ping failed: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL with connection pool")
}
