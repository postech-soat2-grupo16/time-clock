package timeclock

import (
	"gorm.io/gorm"
	"log"
	"time"
	"time-clock/entities"
)

type Repository struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Repository {
	return &Repository{repository: repository}
}

func (r *Repository) ClockIn(userId uint32) (*entities.TimeClock, error) {
	timeClock := entities.TimeClock{
		UserID:    userId,
		ClockIn:   time.Now(),
		CreatedAt: time.Now(),
	}
	result := r.repository.Create(&timeClock)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &timeClock, nil
}
