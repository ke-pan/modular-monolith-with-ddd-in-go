package application

import (
	"context"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"

	"github.com/ke-pan/modular-monolith-with-ddd-in-go/pkg/useraccess/domain"
)

type mockPasswordManager struct {
	hashPassword func(string) string
}

func (m mockPasswordManager) HashPassword(password string) string {
	return m.hashPassword(password)
}

func TestUserRegistrationHandler_RegisterNewUser(t *testing.T) {
	factory := domain.NewUserRegistrationFactory(
		func() string { return "id" },
		func() time.Time { return time.Unix(15600, 0) })
	type fields struct {
		factory         *domain.UserRegistrationFactory
		userCounter     func(string) int
		passwordManager PasswordManager
	}
	type args struct {
		ctx     context.Context
		command *RegisterNewUserCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.UserRegistration
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				factory:         factory,
				userCounter:     func(string) int { return 0 },
				passwordManager: mockPasswordManager{func(s string) string { return "hashed" }},
			},
			args: args{
				ctx: context.Background(),
				command: &RegisterNewUserCommand{
					Login:       "login",
					Password:    "password",
					Email:       "email",
					FirstName:   "john",
					LastName:    "smith",
					ConfirmLink: "link",
				},
			},
			want: domain.UserRegistration{
				ID: "id",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := UserRegistrationHandler{
				factory:         tt.fields.factory,
				userCounter:     tt.fields.userCounter,
				passwordManager: tt.fields.passwordManager,
			}
			got, err := h.RegisterNewUser(tt.args.ctx, tt.args.command)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, got.ID, tt.want.ID)
		})
	}
}
