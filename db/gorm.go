package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	psql "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"

	config "app/config"
)

func InitDB(cfn *config.Config) (*sql.DB, *gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfn.DBHost,
		cfn.DBUser,
		cfn.DBPassword,
		cfn.DBName,
		cfn.DBPort,
	)

	db, err := gorm.Open(psql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Fatalf("Critical: Failed to access underlying SQL pool: %s", err.Error())
	}

	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxIdleTime(time.Hour)

	log.Println("Database connection pool successfully initialized!")
	return sqldb, db
}

// Returns new opned gormdb connection with the provided transcation without creating new connection with postgress.
func GetTxDB(tx *sql.Tx) (*gorm.DB, error) {
	gormDB, err := gorm.Open(psql.New(psql.Config{
		Conn: tx, // Attaches directly to Goose's transaction pipeline
	}), &gorm.Config{})

	return gormDB, err
}
