package main

import (
	config "app/config"
	db "app/db"
	_ "app/db/migrations"
	"time"

	"context"
	"fmt"

	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"
)

func main() {
	// configuration
	var sqlDB *sql.DB
	cnf := config.LoadConfig()
	sqlDB, _ = db.InitDB(cnf)
	defer func() {
			err := sqlDB.Close()
			if err != nil {
				log.Printf("Failed to close Database: %v", err)
			}
		}()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Goose failed to set dialect: %v", err)
	}

	migrationDir, err := filepath.Abs("db/migrations")
	if err != nil {
		log.Fatalf("Could not locate migration directory: %v", err)
	}

	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Printf("Executing database migration command: %v\n", args)

		command := args[0] // up, down, status
		subArgs := args[1:] // any extra flags passed to goose

		// creating a 5 minute safe context to avoid indefinite running of the query
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer func() {
			cancel()
			log.Fatalf("Execution stopped: %v", <-ctx.Done())
		}()

		err := goose.RunContext(ctx, command, sqlDB, migrationDir, subArgs...)
		if err != nil {
			log.Fatalf("Goose command executing failed: %v", err)
		}

		fmt.Println("Migration command executed successfully!")
		return
	}
	fmt.Println("No migration arguments provided. Initializing API server layers...")
}
