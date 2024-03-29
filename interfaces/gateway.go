package interfaces

import (
	"time"
	timeClockAdapter "time-clock/adapters/timeclock"
	"time-clock/entities"
)

type UserGatewayI interface {
	Save(user entities.User) (*entities.User, error)
	GetByRegistration(registration string) (*entities.User, error)
}

type TimeClockGatewayI interface {
	ClockIn(userId uint32) (*entities.TimeClock, error)
	Report(userID uint32, startDate, endDate time.Time) ([]timeClockAdapter.TimeClock, error)
}

type NotificationGatewayI interface {
	ClientSubscriber(user *entities.User) error
	SendNotification(texto string, user *entities.User) error
}
