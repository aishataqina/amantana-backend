package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(b, &a)
}

type PlantCare struct {
	Watering    string `json:"watering"`
	Sunlight    string `json:"sunlight"`
	Temperature string `json:"temperature"`
	Soil        string `json:"soil"`
}

func (pc PlantCare) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

func (pc *PlantCare) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(b, &pc)
}

type Plant struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	Name        string      `json:"name" gorm:"not null"`
	Image       string      `json:"image" gorm:"type:text"`
	Category    string      `json:"category" gorm:"not null"`
	Difficulty  string      `json:"difficulty"`
	Description string      `json:"description" gorm:"type:text"`
	Benefits    StringArray `json:"benefits" gorm:"type:json"`
	Care        PlantCare   `json:"care" gorm:"type:json"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type PlantRepository interface {
	FindAll() ([]Plant, error)
	FindByID(id uint) (*Plant, error)
	Create(plant *Plant) error
}

type PlantUsecase interface {
	GetAll() ([]Plant, error)
	GetByID(id uint) (*Plant, error)
	Create(plant *Plant) error
}
