package user

import (
	"gorm.io/gorm"
	"log"
	"time-clock/entities"
)

type Repository struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Repository {
	return &Repository{repository: repository}
}

func (r *Repository) Save(user entities.User) (*entities.User, error) {
	result := r.repository.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (r *Repository) GetByRegistration(registration string) (*entities.User, error) {
	user := entities.User{
		Registration: registration,
	}
	result := r.repository.Where("registration = ?", registration).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
