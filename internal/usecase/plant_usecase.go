package usecase

import (
	"amantana/internal/domain"
)

type plantUsecase struct {
	plantRepo domain.PlantRepository
}

func NewPlantUsecase(repo domain.PlantRepository) domain.PlantUsecase {
	return &plantUsecase{
		plantRepo: repo,
	}
}

func (u *plantUsecase) GetAll() ([]domain.Plant, error) {
	return u.plantRepo.FindAll()
}

func (u *plantUsecase) GetByID(id uint) (*domain.Plant, error) {
	return u.plantRepo.FindByID(id)
}

func (u *plantUsecase) Create(plant *domain.Plant) error {
	return u.plantRepo.Create(plant)
}
