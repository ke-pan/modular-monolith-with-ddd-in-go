package userregistrastion

import (
	"time"
)

type UserRegistrationID string

type UserRegistrationStatus uint8

const (
	UserRegistrationStatusWaitingForConfirm UserRegistrationStatus = 1
	UserRegistrationStatusConfirmed         UserRegistrationStatus = 2
	UserRegistrationStatusExpired           UserRegistrationStatus = 3
)

type UserRegistration struct {
	ID           UserRegistrationID
	login        string
	password     string
	email        string
	firstName    string
	lastName     string
	name         string
	registerDate time.Time
	status       UserRegistrationStatus
	confirmDate  time.Time
}
