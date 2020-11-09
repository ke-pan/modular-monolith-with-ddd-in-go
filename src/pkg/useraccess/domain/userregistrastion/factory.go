package userregistrastion

import (
	"time"

	"github.com/ke-pan/modular-monolith-with-ddd-in-go/src/pkg/useraccess/domain"
)

type Factory struct {
	genID          func() string
	genCurrentTime func() time.Time
}

var factory *Factory

func NewFactory(genID func() string, genCurrentTime func() time.Time) *Factory {
	if factory != nil {
		return factory
	}
	return &Factory{
		genID:          genID,
		genCurrentTime: genCurrentTime,
	}
}

func (f Factory) RegisterNewUser(login, password, email, firstName, lastName, confirmLink string, countUserWithLogin func(string) int) (UserRegistration, error) {
	if err := domain.CheckRule(UserLoginMustBeUnique(countUserWithLogin, login)); err != nil {
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
