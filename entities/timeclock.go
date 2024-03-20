package entities

import "time"

type TimeClock struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	ClockIn   time.Time `gorm:"not null" json:"clock_in"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
