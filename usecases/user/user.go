package user

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time-clock/entities"
	"time-clock/interfaces"
)

type UseCase struct {
	userGateway interfaces.UserGatewayI
}

func NewUseCase(userGateway interfaces.UserGatewayI) *UseCase {
	return &UseCase{userGateway: userGateway}
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
		log.Fatal(err)
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
