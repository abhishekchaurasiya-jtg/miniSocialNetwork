package main

import (
	// inbuilt imports
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// external imports
	"github.com/pressly/goose/v3"

	// local imports
	config "app/config"
	db "app/db"
	_ "app/db/migrations"
)

func main() {
	sqlDB, _ := db.InitDB(config.LoadConfig())

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Goose failed to set dialect: %s", err.Error())
	}

	migrationDir, err := filepath.Abs("db/migrations")
	if err != nil {
		log.Fatalf("Could not locate migration directory: %s", err.Error())
	}

	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Printf("Executing database migration command: %v\n", args)

		command := args[0]  // up, down, status
		subArgs := args[1:] // any extra flags passed to goose

		err := goose.RunContext(context.Background(), command, sqlDB, migrationDir, subArgs...)
		if err != nil {
			log.Fatalf("Goose command executing failed: %s", err.Error())
		}

		fmt.Println("Migration command executed successfully!")
		return
	}
	fmt.Println("No migration arguments provided. Initializing API server layers...")
}
