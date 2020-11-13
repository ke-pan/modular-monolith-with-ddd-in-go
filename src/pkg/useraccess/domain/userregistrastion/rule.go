package userregistrastion

import (
	"errors"
)

var (
	ErrLoginMustBeUnique                                 = errors.New("user login must be unique")
	ErrUserCannotBeCreatedWhenRegistrationIsNotConfirmed = errors.New("user cannot be created when registration is not conformed")
	ErrUserRegistrationCannotBeConfirmedMoreThanOnce     = errors.New("user registration cannot be confirmed more than once")
	ErrUserRegistrationCannotBeConfirmedAfterExpired     = errors.New("user registration cannot be confirmed after expired")
	ErrUserRegistrationCannotBeExpiredMoreThanOnce       = errors.New("user registration cannot be expired more than once")
)

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

type userRegistrationCannotBeConfirmedMoreThanOnce struct {
	status Status
}

func (r userRegistrationCannotBeConfirmedMoreThanOnce) Validate() error {
	if r.status == StatusConfirmed {
		return ErrUserRegistrationCannotBeConfirmedMoreThanOnce
	}
	return nil
}

func UserRegistrationCannotBeConfirmedMoreThanOnce(status Status) userRegistrationCannotBeConfirmedMoreThanOnce {
	return userRegistrationCannotBeConfirmedMoreThanOnce{status: status}
}

type userRegistrationCannotBeConfirmedAfterExpired struct {
	status Status
}

func (r userRegistrationCannotBeConfirmedAfterExpired) Validate() error {
	if r.status == StatusExpired {
		return ErrUserRegistrationCannotBeConfirmedAfterExpired
	}
	return nil
}

func UserRegistrationCannotBeConfirmedAfterExpired(status Status) userRegistrationCannotBeConfirmedAfterExpired {
	return userRegistrationCannotBeConfirmedAfterExpired{status: status}
}

type userRegistrationCannotBeExpiredMoreThanOnce struct {
	status Status
}

func (r userRegistrationCannotBeExpiredMoreThanOnce) Validate() error {
	if r.status == StatusExpired {
		return ErrUserRegistrationCannotBeExpiredMoreThanOnce
	}
	return nil
}

func UserRegistrationCannotBeExpiredMoreThanOnce(status Status) userRegistrationCannotBeExpiredMoreThanOnce {
	return userRegistrationCannotBeExpiredMoreThanOnce{status: status}
}
