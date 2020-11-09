package rule

import "errors"

var ErrLoginMustBeUnique = errors.New("user login must be unique")

type Rule interface {
	Validate() error
}

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
