package postgresql_db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() *sql.DB {
	dsn := "host=localhost user=root password=11111111 dbname=cyberus_db port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// Get generic database object sql.DB to configure pool
	PostgreSqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	// ✅ Configure connection pool
	PostgreSqlDB.SetMaxOpenConns(100)                 // max open connections
	PostgreSqlDB.SetMaxIdleConns(5)                   // max idle connections
	PostgreSqlDB.SetConnMaxLifetime(30 * time.Second) // max lifetime of a connection

	// Test connection
	err = PostgreSqlDB.Ping()
	if err != nil {
		log.Fatalf("ping failed: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL with connection pool")

	return PostgreSqlDB
}
