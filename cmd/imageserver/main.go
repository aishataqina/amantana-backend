package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Serve static files from assets directory
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Add endpoint untuk mengecek keberadaan gambar
	http.HandleFunc("/check-images", func(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("Server berjalan di http://localhost:3000")
	fmt.Println("Buka http://localhost:3000/check-images untuk melihat daftar gambar")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
