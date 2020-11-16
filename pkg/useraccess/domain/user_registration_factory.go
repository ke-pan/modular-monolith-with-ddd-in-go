package domain

import (
	"time"
)

type UserRegistrationFactory struct {
	genID          func() string
	genCurrentTime func() time.Time
}

var factory *UserRegistrationFactory

func NewUserRegistrationFactory(genID func() string, genCurrentTime func() time.Time) *UserRegistrationFactory {
	if factory != nil {
		return factory
	}
	return &UserRegistrationFactory{
		genID:          genID,
		genCurrentTime: genCurrentTime,
	}
}

func (f UserRegistrationFactory) RegisterNewUser(login, password, email, firstName, lastName, confirmLink string, countUserWithLogin func(string) int) (UserRegistration, error) {
	if err := CheckRule(UserLoginMustBeUnique(countUserWithLogin, login)); err != nil {
		return UserRegistration{}, err
	}
	return UserRegistration{
		ID:           ID(f.genID()),
		login:        login,
		password:     password,
		email:        email,
		firstName:    firstName,
		lastName:     lastName,
		name:         firstName + " " + lastName,
		registerDate: f.genCurrentTime(),
		status:       StatusWaitingForConfirm,
		confirmDate:  f.genCurrentTime(),
	}, nil
}
