package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	// Serve static files dari direktori uploads
	workDir, _ := os.Getwd()
	uploadsDir := http.Dir(filepath.Join(workDir, "uploads"))
	FileServer(r, "/uploads", uploadsDir)

	// Register handlers
	plantHttp.NewPlantHandler(r, plantUsecase)

	// Start server
	port := viper.GetInt("app.port")
	fmt.Printf("Server starting on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

// FileServer adalah helper untuk serving static files
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
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
