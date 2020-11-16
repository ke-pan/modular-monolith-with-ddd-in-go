package domain

import (
	"time"
)

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

func (ur UserRegistration) CreateUser() (User, error) {
	if err := CheckRule(UserCannotBeCreatedWhenRegistrationIsNotConfirmed(ur.status)); err != nil {
		return User{}, err
	}
	return CreateFromUserRegistration(ur.ID, ur.login, ur.password, ur.email, ur.firstName, ur.lastName, ur.name)
}

func (ur *UserRegistration) Confirm() error {
	if err := CheckRule(UserRegistrationCannotBeConfirmedMoreThanOnce(ur.status)); err != nil {
		return err
	}
	if err := CheckRule(UserRegistrationCannotBeConfirmedAfterExpired(ur.status)); err != nil {
		return err
	}
	ur.status = StatusConfirmed
	ur.confirmDate = time.Now()
	return nil
}

func (ur *UserRegistration) Expire() error {
	if err := CheckRule(UserRegistrationCannotBeExpiredMoreThanOnce(ur.status)); err != nil {
		return err
	}
	ur.status = StatusExpired
	return nil
}
