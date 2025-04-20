package postgresql_db

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgreSqlInstance() (*gorm.DB, *sql.DB, error) {
	dsn := "host=localhost user=root password=11111111 dbname=cyberus_db port=5432 sslmode=disable TimeZone=Asia/Bangkok search_path=root@cyberus"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// Get generic database object sql.DB to configure pool
	sqlConfig, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	// ✅ Configure connection pool
	sqlConfig.SetMaxOpenConns(100)                 // max open connections
	sqlConfig.SetMaxIdleConns(5)                   // max idle connections
	sqlConfig.SetConnMaxLifetime(30 * time.Second) // max lifetime of a connection

	return db, sqlConfig, err
}
