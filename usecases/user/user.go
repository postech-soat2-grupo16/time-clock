package user

import (
	"errors"
	"gorm.io/gorm"
	timeClockAdapter "time-clock/adapters/timeclock"
	"time-clock/entities"
	"time-clock/interfaces"
)

type UseCase struct {
	userGateway      interfaces.UserGatewayI
	timeClockGateway interfaces.TimeClockGatewayI
}

func NewUseCase(userGateway interfaces.UserGatewayI, timeClockGateway interfaces.TimeClockGatewayI) *UseCase {
	return &UseCase{
		userGateway:      userGateway,
		timeClockGateway: timeClockGateway,
	}
}

func (u *UseCase) Create(name, email, registration, password string) (*entities.User, error) {
	user := entities.User{
		Name:         name,
		Email:        email,
		Registration: registration,
		Password:     password,
	}

	result, err := u.userGateway.Save(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UseCase) GetByRegistration(registration string) (*entities.User, error) {
	result, err := u.userGateway.GetByRegistration(registration)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

func (u *UseCase) ClockIn(registration string) (*entities.TimeClock, error) {
	user, err := u.GetByRegistration(registration)
	if err != nil {
		return nil, err
	}

	result, err := u.timeClockGateway.ClockIn(user.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

func (u *UseCase) Report(userID, year, month, day uint32) ([]timeClockAdapter.TimeClock, error) {
	result, err := u.timeClockGateway.Report(userID, year, month, day)
	if err != nil {
		return nil, err
	}

	return result, nil
}
