package user

type ID string

type Role uint8

const (
	RoleMember Role = 1
)

type User struct {
	ID        ID
	login     string
	password  string
	email     string
	active    bool
	firstName string
	lastName  string
	name      string
	roles     []Role
}
