package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"amantana/internal/domain"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type PlantHandler struct {
	plantUsecase domain.PlantUsecase
}

func NewPlantHandler(r *chi.Mux, plantUsecase domain.PlantUsecase) {
	handler := &PlantHandler{
		plantUsecase: plantUsecase,
	}

	r.Get("/plants", handler.GetPlants)
	r.Get("/plants/{id}", handler.GetPlantByID)
	r.Post("/plants", handler.CreatePlant)
}

func (h *PlantHandler) GetPlants(w http.ResponseWriter, r *http.Request) {
	plants, err := h.plantUsecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, plants)
}

func (h *PlantHandler) GetPlantByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	plant, err := h.plantUsecase.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, plant)
}

func (h *PlantHandler) CreatePlant(w http.ResponseWriter, r *http.Request) {
	var plant domain.Plant
	contentType := r.Header.Get("Content-Type")

	fmt.Printf("Received request with Content-Type: %s\n", contentType)

	// Handle multipart form data
	if strings.HasPrefix(contentType, "multipart/form-data") {
		fmt.Println("Processing multipart form data")

		// Parse multipart form dengan maksimal ukuran 10MB
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "Ukuran file terlalu besar, maksimal 10MB", http.StatusBadRequest)
			return
		}

		// Handle file upload
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error mengambil file image: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Validasi tipe file
		fileContentType := handler.Header.Get("Content-Type")
		fmt.Printf("File Content-Type: %s\n", fileContentType)

		if !isValidImageType(fileContentType) {
			http.Error(w, "Tipe file tidak valid. Hanya menerima: jpg, jpeg, png", http.StatusBadRequest)
			return
		}

		// Generate unique filename
		filename := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(handler.Filename))
		fmt.Printf("Generated filename: %s\n", filename)

		// Pastikan direktori uploads ada
		if err := os.MkdirAll("uploads", 0755); err != nil {
			http.Error(w, "Error membuat direktori uploads: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Simpan file
		filepath := filepath.Join("uploads", filename)
		dst, err := os.Create(filepath)
		if err != nil {
			http.Error(w, "Error menyimpan file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			os.Remove(filepath)
			http.Error(w, "Error menyalin file: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("File uploaded successfully")

		// Set data plant dari form
		plant = domain.Plant{
			Name:        r.FormValue("name"),
			Category:    r.FormValue("category"),
			Difficulty:  r.FormValue("difficulty"),
			Description: r.FormValue("description"),
			Image:       filepath,
		}

		fmt.Printf("Form data received: %s\n", debugJSON(plant))

		// Parse benefits
		if benefits := r.FormValue("benefits"); benefits != "" {
			fmt.Printf("Processing benefits: %s\n", benefits)
			var benefitsArray []string
			if err := json.Unmarshal([]byte(benefits), &benefitsArray); err != nil {
				fmt.Printf("Failed to parse benefits as array, using as single string: %v\n", err)
				plant.Benefits = domain.StringArray{benefits}
			} else {
				plant.Benefits = benefitsArray
			}
		}

		// Parse care
		if care := r.FormValue("care"); care != "" {
			fmt.Printf("Processing care: %s\n", care)
			if err := json.Unmarshal([]byte(care), &plant.Care); err != nil {
				fmt.Printf("Failed to parse care as JSON, using form fields: %v\n", err)
				// Jika parsing JSON gagal, gunakan form fields
				plant.Care = domain.PlantCare{
					Watering:    r.FormValue("care_watering"),
					Sunlight:    r.FormValue("care_sunlight"),
					Temperature: r.FormValue("care_temperature"),
					Soil:        r.FormValue("care_soil"),
				}
			}
		} else {
			// Gunakan form fields langsung
			plant.Care = domain.PlantCare{
				Watering:    r.FormValue("care_watering"),
				Sunlight:    r.FormValue("care_sunlight"),
				Temperature: r.FormValue("care_temperature"),
				Soil:        r.FormValue("care_soil"),
			}
		}
	} else {
		fmt.Println("Processing JSON request")
		// Handle JSON request
		if err := json.NewDecoder(r.Body).Decode(&plant); err != nil {
			http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Validasi required fields
	if plant.Name == "" || plant.Category == "" {
		http.Error(w, "Name dan category harus diisi", http.StatusBadRequest)
		return
	}

	fmt.Printf("Final plant data before save: %s\n", debugJSON(plant))

	// Simpan ke database
	if err := h.plantUsecase.Create(&plant); err != nil {
		if plant.Image != "" {
			os.Remove(plant.Image)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, plant)
}

// Helper function untuk validasi tipe file
func isValidImageType(contentType string) bool {
	validTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
	}
	return validTypes[contentType]
}
