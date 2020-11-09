package userregistrastion

import (
	"testing"
	"time"

	"github.com/ke-pan/modular-monolith-with-ddd-in-go/src/pkg/useraccess/domain/user"
	"github.com/magiconair/properties/assert"
)

func TestUserRegistration_CreateUser(t *testing.T) {
	type fields struct {
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
	var tests = []struct {
		name       string
		fields     fields
		wantUserID user.ID
		wantErr    error
	}{
		{
			name: "success",
			fields: fields{
				ID:           "id",
				login:        "john",
				password:     "passwd",
				email:        "email",
				firstName:    "john",
				lastName:     "smith",
				name:         "john smith",
				registerDate: time.Unix(156000, 0),
				status:       StatusConfirmed,
				confirmDate:  time.Unix(156000, 0),
			},
			wantUserID: "id",
			wantErr:    nil,
		},
		{
			name: "status is not confirmed",
			fields: fields{
				ID:           "id",
				login:        "john",
				password:     "passwd",
				email:        "email",
				firstName:    "john",
				lastName:     "smith",
				name:         "john smith",
				registerDate: time.Unix(156000, 0),
				status:       StatusWaitingForConfirm,
				confirmDate:  time.Unix(156000, 0),
			},
			wantUserID: "",
			wantErr:    ErrUserCannotBeCreatedWhenRegistrationIsNotConfirmed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := UserRegistration{
				ID:           tt.fields.ID,
				login:        tt.fields.login,
				password:     tt.fields.password,
				email:        tt.fields.email,
				firstName:    tt.fields.firstName,
				lastName:     tt.fields.lastName,
				name:         tt.fields.name,
				registerDate: tt.fields.registerDate,
				status:       tt.fields.status,
				confirmDate:  tt.fields.confirmDate,
			}
			got, err := ur.CreateUser()
			if tt.wantErr != nil {
				assert.Equal(t, err, tt.wantErr)
			} else {
				assert.Equal(t, got.ID, tt.wantUserID)
			}
		})
	}
}
