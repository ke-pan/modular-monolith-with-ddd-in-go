package domain

func CreateFromUserRegistration(id ID, login, password, email, firstName, lastName, name string) (User, error) {
	return User{
		ID:        id,
		login:     login,
		password:  password,
		email:     email,
		active:    true,
		firstName: firstName,
		lastName:  lastName,
		name:      name,
		roles:     []Role{RoleMember},
	}, nil
}
