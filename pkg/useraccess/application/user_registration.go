package application

import (
	"context"

	"github.com/ke-pan/modular-monolith-with-ddd-in-go/pkg/useraccess/domain"
)

type UserRegistrationRepository interface {
	GetByID(domain.ID) (domain.UserRegistration, error)
	Add(registration domain.UserRegistration) error
}

type PasswordManager interface {
	HashPassword(string) string
}

type UserRegistrationHandler struct {
	repo            UserRegistrationRepository
	factory         *domain.UserRegistrationFactory
	userCounter     func(string) int
	passwordManager PasswordManager
}

func NewUserRegistrationHandler(repo UserRegistrationRepository, factory *domain.UserRegistrationFactory, manager PasswordManager, counter func(string) int) *UserRegistrationHandler {
	return &UserRegistrationHandler{repo: repo, factory: factory, passwordManager: manager, userCounter: counter}
}

type RegisterNewUserCommand struct {
	Login       string
	Password    string
	Email       string
	FirstName   string
	LastName    string
	ConfirmLink string
}

func (h UserRegistrationHandler) RegisterNewUser(ctx context.Context, command *RegisterNewUserCommand) (domain.UserRegistration, error) {
	password := h.passwordManager.HashPassword(command.Password)

	ur, err := h.factory.RegisterNewUser(command.Login, password, command.Email, command.FirstName, command.LastName, command.ConfirmLink, h.userCounter)
	if err != nil {
		return domain.UserRegistration{}, err
	}

	if err := h.repo.Add(ur); err != nil {
		return domain.UserRegistration{}, err
	}

	return ur, nil
}
