package postgres

import (
	"amantana/internal/domain"

	"gorm.io/gorm"
)

type plantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) domain.PlantRepository {
	return &plantRepository{
		db: db,
	}
}

func (r *plantRepository) FindAll() ([]domain.Plant, error) {
	var plants []domain.Plant
	err := r.db.Find(&plants).Error
	return plants, err
}

func (r *plantRepository) FindByID(id uint) (*domain.Plant, error) {
	var plant domain.Plant
	err := r.db.First(&plant, id).Error
	if err != nil {
		return nil, err
	}
	return &plant, nil
}

func (r *plantRepository) Create(plant *domain.Plant) error {
	return r.db.Create(plant).Error
}
