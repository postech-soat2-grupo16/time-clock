package entities

import (
	"time"
)

type User struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Email        string    `gorm:"not null" json:"email"`
	Registration string    `gorm:"not null" json:"registration"`
	Password     string    `gorm:"not null" json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
