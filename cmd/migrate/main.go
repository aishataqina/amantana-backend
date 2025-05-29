package main

import (
	"fmt"
	"log"
	"os"

	"amantana/internal/domain"
	plantRepo "amantana/internal/repository/postgres"

	"github.com/spf13/viper"
	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate/main.go [migrate|seed|rollback]")
	}

	command := os.Args[1]

	// Load configuration
	loadConfig()

	// Initialize database
	db := initDB()

	switch command {
	case "migrate":
		migrate(db)
	case "seed":
		seed(db)
	case "rollback":
		rollback(db)
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

func loadConfig() {
	viper.SetConfigFile("configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.ssl_mode"),
	)

	db, err := gorm.Open(pgDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	return db
}

func migrate(db *gorm.DB) {
	fmt.Println("Running migrations...")
	if err := db.AutoMigrate(&domain.Plant{}); err != nil {
		log.Fatalf("Error running migrations: %s", err)
	}
	fmt.Println("Migrations completed successfully!")
}

func seed(db *gorm.DB) {
	fmt.Println("Running seeds...")
	if err := plantRepo.SeedPlants(db); err != nil {
		log.Fatalf("Error running seeds: %s", err)
	}
	fmt.Println("Seeds completed successfully!")
}

func rollback(db *gorm.DB) {
	fmt.Println("Rolling back migrations...")
	if err := db.Migrator().DropTable(&domain.Plant{}); err != nil {
		log.Fatalf("Error rolling back migrations: %s", err)
	}
	fmt.Println("Rollback completed successfully!")
}
