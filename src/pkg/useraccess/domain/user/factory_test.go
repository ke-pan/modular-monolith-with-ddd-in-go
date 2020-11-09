package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFromUserRegistration(t *testing.T) {
	type args struct {
		id        ID
		login     string
		password  string
		email     string
		firstName string
		lastName  string
		name      string
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id:        "id",
				login:     "login",
				password:  "passwd",
				email:     "email",
				firstName: "firstName",
				lastName:  "lastName",
				name:      "name",
			},
			want: User{
				ID:        "id",
				login:     "login",
				password:  "passwd",
				email:     "email",
				active:    true,
				firstName: "firstName",
				lastName:  "lastName",
				name:      "name",
				roles:     []Role{RoleMember},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFromUserRegistration(tt.args.id, tt.args.login, tt.args.password, tt.args.email, tt.args.firstName, tt.args.lastName, tt.args.name)
			if tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}
