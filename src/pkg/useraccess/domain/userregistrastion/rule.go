package userregistrastion

import (
	"errors"
)

var ErrLoginMustBeUnique = errors.New("user login must be unique")
var ErrUserCannotBeCreatedWhenRegistrationIsNotConfirmed = errors.New("user cannot be created when registration is not conformed")

type userLoginMustBeUnique struct {
	countUserWithLogin func() int
}

// Validate implements Rule interface
func (r userLoginMustBeUnique) Validate() error {
	if r.countUserWithLogin() > 0 {
		return ErrLoginMustBeUnique
	}
	return nil
}

func UserLoginMustBeUnique(countUserWithLogin func(login string) int, login string) userLoginMustBeUnique {
	return userLoginMustBeUnique{
		countUserWithLogin: func() int {
			return countUserWithLogin(login)
		},
	}
}

type userCannotBeCreatedWhenRegistrationIsNotConfirmed struct {
	status Status
}

func (r userCannotBeCreatedWhenRegistrationIsNotConfirmed) Validate() error {
	if r.status != StatusConfirmed {
		return ErrUserCannotBeCreatedWhenRegistrationIsNotConfirmed
	}
	return nil
}

func UserCannotBeCreatedWhenRegistrationIsNotConfirmed(status Status) userCannotBeCreatedWhenRegistrationIsNotConfirmed {
	return userCannotBeCreatedWhenRegistrationIsNotConfirmed{
		status: status,
	}
}
