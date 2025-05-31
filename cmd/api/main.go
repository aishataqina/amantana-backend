package main

import (
	"fmt"
	"io"
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

func moveUploadsToAssets() error {
	// Cek apakah direktori uploads ada
	uploadsDir := "uploads"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		fmt.Println("Direktori uploads tidak ditemukan")
		return nil
	}

	// Pastikan direktori assets/images ada
	assetsDir := "assets/images"
	if err := os.MkdirAll(assetsDir, 0755); err != nil {
		return fmt.Errorf("gagal membuat direktori assets/images: %v", err)
	}

	// Baca semua file di direktori uploads
	files, err := os.ReadDir(uploadsDir)
	if err != nil {
		return fmt.Errorf("gagal membaca direktori uploads: %v", err)
	}

	// Pindahkan setiap file
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		sourcePath := filepath.Join(uploadsDir, file.Name())
		destPath := filepath.Join(assetsDir, file.Name())

		// Buka file sumber
		source, err := os.Open(sourcePath)
		if err != nil {
			fmt.Printf("Gagal membuka file %s: %v\n", file.Name(), err)
			continue
		}
		defer source.Close()

		// Buat file tujuan
		destination, err := os.Create(destPath)
		if err != nil {
			fmt.Printf("Gagal membuat file %s: %v\n", file.Name(), err)
			continue
		}
		defer destination.Close()

		// Salin konten file
		if _, err := io.Copy(destination, source); err != nil {
			fmt.Printf("Gagal menyalin file %s: %v\n", file.Name(), err)
			continue
		}

		// Hapus file dari uploads setelah berhasil disalin
		if err := os.Remove(sourcePath); err != nil {
			fmt.Printf("Gagal menghapus file %s dari uploads: %v\n", file.Name(), err)
		}

		fmt.Printf("Berhasil memindahkan %s ke assets/images\n", file.Name())
	}

	return nil
}

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

	// Pindahkan file dari uploads ke assets/images
	if err := moveUploadsToAssets(); err != nil {
		log.Printf("Error saat memindahkan file: %v\n", err)
	}

	// Serve static files dari direktori uploads dan assets
	workDir, _ := os.Getwd()
	uploadsDir := http.Dir(filepath.Join(workDir, "uploads"))
	assetsDir := http.Dir(filepath.Join(workDir, "assets"))
	FileServer(r, "/uploads", uploadsDir)
	FileServer(r, "/assets", assetsDir)

	// Tambahkan endpoint untuk mengecek gambar
	r.Get("/check-images", func(w http.ResponseWriter, r *http.Request) {
		imagesDir := "assets/images"
		files, err := os.ReadDir(imagesDir)
		if err != nil {
			http.Error(w, "Gagal membaca direktori images", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "<h1>Daftar Gambar yang Ada:</h1>")
		fmt.Fprintf(w, "<ul>")
		for _, file := range files {
			if !file.IsDir() {
				imagePath := filepath.Join("/assets/images", file.Name())
				fmt.Fprintf(w, "<li><a href='%s'>%s</a></li>", imagePath, file.Name())
			}
		}
		fmt.Fprintf(w, "</ul>")

		// Tambahkan daftar gambar yang seharusnya ada berdasarkan seeder
		expectedImages := []string{
			"philodendron-pink-princess.jpg",
			"alocasia-black-velvet.jpg",
			"lidah-buaya.jpg",
			"bird-of-paradise.jpg",
			"monstera-deliciosa.jpg",
			"sirih-gading.jpg",
			"kangkung.jpeg",
			"tomat.jpg",
			"kemangi.jpg",
			"chinese-evergreen.jpg",
			"kaktus-mini.jpeg",
			"jahe.jpg",
			"lavender.jpg",
			"bayam.jpg",
			"basil.jpg",
			"cabe-rawit.jpg",
			"fiddle-leaf-fig.jpg",
			"jasmine.jpg",
			"pachira-aquatica.jpg",
			"mint.jpg",
			"bonsai-serut.jpg",
			"anggrek-bulan.jpg",
			"bayam-merah.jpg",
			"bunga-matahari.jpg",
		}

		fmt.Fprintf(w, "<h2>Gambar yang Belum Ada:</h2>")
		fmt.Fprintf(w, "<ul style='color: red;'>")
		for _, expectedImage := range expectedImages {
			if _, err := os.Stat(filepath.Join(imagesDir, expectedImage)); os.IsNotExist(err) {
				fmt.Fprintf(w, "<li>%s</li>", expectedImage)
			}
		}
		fmt.Fprintf(w, "</ul>")
	})

	// Register handlers
	plantHttp.NewPlantHandler(r, plantUsecase)

	// Start server
	port := viper.GetInt("app.port")
	fmt.Printf("Server starting on port %d...\n", port)
	fmt.Printf("Buka http://localhost:%d/check-images untuk melihat daftar gambar\n", port)
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
