package interfaces

import (
	timeClockAdapter "time-clock/adapters/timeclock"
	"time-clock/entities"
)

type UserGatewayI interface {
	Save(user entities.User) (*entities.User, error)
	GetByRegistration(registration string) (*entities.User, error)
}

type TimeClockGatewayI interface {
	ClockIn(userId uint32) (*entities.TimeClock, error)
	Report(userID, year, month, day uint32) ([]timeClockAdapter.TimeClock, error)
}
