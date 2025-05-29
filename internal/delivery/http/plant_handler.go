package http

import (
	"net/http"
	"strconv"

	"amantana/internal/domain"

	"github.com/go-chi/chi/v5"
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
	if err := parseJSON(r, &plant); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.plantUsecase.Create(&plant); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, plant)
}
