package user

import (
	"errors"
	"fmt"
	timeClockAdapter "time-clock/adapters/timeclock"
	"time-clock/entities"
	"time-clock/interfaces"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UseCase struct {
	userGateway         interfaces.UserGatewayI
	timeClockGateway    interfaces.TimeClockGatewayI
	notificationGateway interfaces.NotificationGatewayI
}

func NewUseCase(userGateway interfaces.UserGatewayI,
	timeClockGateway interfaces.TimeClockGatewayI,
	notificationGateway interfaces.NotificationGatewayI) *UseCase {
	return &UseCase{
		userGateway:         userGateway,
		timeClockGateway:    timeClockGateway,
		notificationGateway: notificationGateway,
	}
}

func (u *UseCase) Create(name, email, registration, password string) (*entities.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := entities.User{
		Name:         name,
		Email:        email,
		Registration: registration,
		Password:     string(hash),
	}

	result, err := u.userGateway.Save(user)
	if err != nil {
		return nil, err
	}

	err = u.notificationGateway.ClientSubscriber(result)
	if err != nil {
		fmt.Printf("error subscribing")
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
