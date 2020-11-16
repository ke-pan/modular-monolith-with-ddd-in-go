package domain

import (
	"testing"
	"time"

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
		wantUserID ID
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

func TestUserRegistration_Confirm(t *testing.T) {
	type fields struct {
		status Status
	}
	tests := []struct {
		name   string
		fields fields
		want   Status
		err    error
	}{
		{
			name:   "success",
			fields: fields{status: StatusWaitingForConfirm},
			want:   StatusConfirmed,
			err:    nil,
		},
		{
			name:   "registration can not be confirmed more than once",
			fields: fields{status: StatusConfirmed},
			want:   StatusConfirmed,
			err:    ErrUserRegistrationCannotBeConfirmedMoreThanOnce,
		},
		{
			name:   "registration can not be confirmed after expired",
			fields: fields{status: StatusExpired},
			want:   StatusExpired,
			err:    ErrUserRegistrationCannotBeConfirmedAfterExpired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRegistration{
				status: tt.fields.status,
			}
			err := ur.Confirm()
			assert.Equal(t, err, tt.err)
			assert.Equal(t, ur.status, tt.want)
		})
	}
}

func TestUserRegistration_Expire(t *testing.T) {
	type fields struct {
		status Status
	}
	tests := []struct {
		name    string
		fields  fields
		want    Status
		wantErr error
	}{
		{
			name:    "success",
			fields:  fields{status: StatusConfirmed},
			want:    StatusExpired,
			wantErr: nil,
		},
		{
			name:    "registration cannot be expired more than once",
			fields:  fields{status: StatusExpired},
			want:    StatusExpired,
			wantErr: ErrUserRegistrationCannotBeExpiredMoreThanOnce,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRegistration{
				status: tt.fields.status,
			}
			err := ur.Expire()
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, ur.status, tt.want)
		})
	}
}
