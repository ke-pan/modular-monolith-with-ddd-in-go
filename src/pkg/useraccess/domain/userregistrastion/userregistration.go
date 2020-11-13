package userregistrastion

import (
	"time"

	"github.com/ke-pan/modular-monolith-with-ddd-in-go/src/pkg/useraccess/domain"
	"github.com/ke-pan/modular-monolith-with-ddd-in-go/src/pkg/useraccess/domain/user"
)

type ID string

type Status uint8

const (
	StatusWaitingForConfirm Status = 1
	StatusConfirmed         Status = 2
	StatusExpired           Status = 3
)

type UserRegistration struct {
	ID           ID
	login        string
	password     string
	email        string
	firstName    string
	lastName     string
	name         string
	registerDate time.Time
	status       Status
	confirmDate  time.Time
}

func (ur UserRegistration) CreateUser() (user.User, error) {
	if err := domain.CheckRule(UserCannotBeCreatedWhenRegistrationIsNotConfirmed(ur.status)); err != nil {
		return user.User{}, err
	}
	return user.CreateFromUserRegistration(user.ID(ur.ID), ur.login, ur.password, ur.email, ur.firstName, ur.lastName, ur.name)
}

func (ur *UserRegistration) Confirm() error {
	if err := domain.CheckRule(UserRegistrationCannotBeConfirmedMoreThanOnce(ur.status)); err != nil {
		return err
	}
	if err := domain.CheckRule(UserRegistrationCannotBeConfirmedAfterExpired(ur.status)); err != nil {
		return err
	}
	ur.status = StatusConfirmed
	ur.confirmDate = time.Now()
	return nil
}

func (ur *UserRegistration) Expire() error {
	if err := domain.CheckRule(UserRegistrationCannotBeExpiredMoreThanOnce(ur.status)); err != nil {
		return err
	}
	ur.status = StatusExpired
	return nil
}
