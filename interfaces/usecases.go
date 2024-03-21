package interfaces

import (
	"time"
	timeClockAdapter "time-clock/adapters/timeclock"
	"time-clock/entities"
)

type UserUserCase interface {
	Create(name, email, registration, password string) (*entities.User, error)
	GetByRegistration(registration string) (*entities.User, error)
	ClockIn(registration string) (*entities.TimeClock, error)
	Report(userID uint32, startDate, endDate time.Time) ([]timeClockAdapter.TimeClock, error)
}
