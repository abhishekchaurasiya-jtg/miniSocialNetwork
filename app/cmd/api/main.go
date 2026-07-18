package main

import (
	config "app/config"
	db_pool "app/db"
	_ "app/db/migrations"
	
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"
)

func main() {
	cnf := config.LoadConfig()
	sqlDB := db_pool.InitDB(cnf)
	defer sqlDB.Close()

	fmt.Println("Db Connected..")

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Goose failed to set dialect: %v", err)
	}

	migrationDir, err := resolveMigrationDir()
	if err != nil {
		log.Fatalf("Could not locate migration directory: %v", err)
	}

	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Printf("Executing database migration command: %v\n", args)

		command := args[0]    // e.g., "up", "down", "status"
		subArgs := args[1:]    // any extra flags passed to goose

		err := goose.RunContext(context.Background(), command, sqlDB, migrationDir, subArgs...)
		if err != nil {
			log.Fatalf("Goose command execution failed: %v", err)
		}

		fmt.Println("Migration command executed successfully!")
		return
	}

	// -----------------------------------------------------------------
	// 4. Future Gin Web Server Entrypoint
	// -----------------------------------------------------------------
	fmt.Println("No migration arguments provided. Initializing API server layers...")
	// r := gin.Default()
	// r.Run(":8080")
}

func resolveMigrationDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	repoRoot := cwd
	for {
		if _, err := os.Stat(filepath.Join(repoRoot, "go.mod")); err == nil {
			break
		}

		parent := filepath.Dir(repoRoot)
		if parent == repoRoot {
			return "", fmt.Errorf("could not find repository root")
		}
		repoRoot = parent
	}

	candidates := []string{
		filepath.Join(repoRoot, "db", "migrations"),
		filepath.Join(repoRoot, "app", "db", "migrations"),
	}

	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("migration directory not found; checked %v", candidates)
}
