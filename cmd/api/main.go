package main

import (
	"fmt"
	"log"
	"net/http"

	plantHttp "amantana/internal/delivery/http"
	plantRepo "amantana/internal/repository/postgres"
	"amantana/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	loadConfig()

	// Initialize database
	db := initDB()

	// Initialize repository
	plantRepository := plantRepo.NewPlantRepository(db)

	// Initialize usecase
	plantUsecase := usecase.NewPlantUsecase(plantRepository)

	// Initialize router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Register handlers
	plantHttp.NewPlantHandler(r, plantUsecase)

	// Start server
	port := viper.GetInt("app.port")
	fmt.Printf("Server starting on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
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
