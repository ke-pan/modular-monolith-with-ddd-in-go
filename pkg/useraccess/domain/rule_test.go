package domain

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestUserLoginMustBeUnique(t *testing.T) {
	type args struct {
		countUserWithLogin func(login string) int
		login              string
	}
	tests := []struct {
		name        string
		args        args
		validateErr error
	}{
		{
			name: "login is unique",
			args: args{
				countUserWithLogin: func(login string) int {
					return 0
				},
				login: "john",
			},
			validateErr: nil,
		},
		{
			name: "login is not unique",
			args: args{
				countUserWithLogin: func(login string) int {
					return 1
				},
				login: "john",
			},
			validateErr: ErrLoginMustBeUnique,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := UserLoginMustBeUnique(tt.args.countUserWithLogin, tt.args.login)
			assert.Equal(t, rule.Validate(), tt.validateErr)
		})
	}
}
