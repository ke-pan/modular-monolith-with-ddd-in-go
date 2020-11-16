package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFactory_RegisterNewUser(t *testing.T) {
	type fields struct {
		genID          func() string
		genCurrentTime func() time.Time
	}
	type args struct {
		login              string
		password           string
		email              string
		firstName          string
		lastName           string
		confirmLink        string
		countUserWithLogin func(string) int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    UserRegistration
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				genID: func() string {
					return "id"
				},
				genCurrentTime: func() time.Time {
					return time.Unix(1560000, 0)
				},
			},
			args: args{
				login:       "john",
				password:    "2df92ndkuhy",
				email:       "john@email.com",
				firstName:   "john",
				lastName:    "smith",
				confirmLink: "https://asdhfin",
				countUserWithLogin: func(s string) int {
					return 0
				},
			},
			want: UserRegistration{
				ID:           ID("id"),
				login:        "john",
				password:     "2df92ndkuhy",
				email:        "john@email.com",
				firstName:    "john",
				lastName:     "smith",
				name:         "john smith",
				registerDate: time.Unix(1560000, 0),
				status:       StatusWaitingForConfirm,
				confirmDate:  time.Unix(1560000, 0),
			},
			wantErr: false,
		},
		{
			name: "rule UserLoginMustBeUnique broke",
			fields: fields{
				genID: func() string {
					return "id"
				},
				genCurrentTime: func() time.Time {
					return time.Unix(1560000, 0)
				},
			},
			args: args{
				login:       "john",
				password:    "2df92ndkuhy",
				email:       "john@email.com",
				firstName:   "john",
				lastName:    "smith",
				confirmLink: "https://asdhfin",
				countUserWithLogin: func(s string) int {
					return 1
				},
			},
			want:    UserRegistration{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := UserRegistrationFactory{
				genID:          tt.fields.genID,
				genCurrentTime: tt.fields.genCurrentTime,
			}
			got, err := f.RegisterNewUser(tt.args.login, tt.args.password, tt.args.email, tt.args.firstName, tt.args.lastName, tt.args.confirmLink, tt.args.countUserWithLogin)
			if tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}
