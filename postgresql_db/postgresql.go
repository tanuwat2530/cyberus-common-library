package postgresql_db

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgreSqlInstance(connectionStr string) (*gorm.DB, *sql.DB, error) {

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
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
