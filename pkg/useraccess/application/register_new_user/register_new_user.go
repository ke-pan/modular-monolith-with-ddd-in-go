package register_new_user

import (
	"context"

	"github.com/ke-pan/modular-monolith-with-ddd-in-go/pkg/useraccess/domain"
)

type UserRegistrationRepository interface {
	Insert(registration domain.UserRegistration) error
}

type PasswordManager interface {
	HashPassword(string) string
}

type UserRegisteredHandler interface {
	Handle(UserRegisteredEvent) error
}

type UserRegisteredEvent struct {
	domain.UserRegistration
}

type Handler struct {
	repo                  UserRegistrationRepository
	factory               *domain.UserRegistrationFactory
	userCounter           func(string) int
	passwordManager       PasswordManager
	userRegisteredHandler UserRegisteredHandler
}

func NewHandler(repo UserRegistrationRepository, factory *domain.UserRegistrationFactory, manager PasswordManager, counter func(string) int, eventHandler UserRegisteredHandler) *Handler {
	return &Handler{repo: repo, factory: factory, passwordManager: manager, userCounter: counter, userRegisteredHandler: eventHandler}
}

type RegisterNewUserCommand struct {
	Login       string
	Password    string
	Email       string
	FirstName   string
	LastName    string
	ConfirmLink string
}

func (h Handler) Handle(ctx context.Context, command *RegisterNewUserCommand) (domain.UserRegistration, error) {
	password := h.passwordManager.HashPassword(command.Password)

	ur, err := h.factory.RegisterNewUser(command.Login, password, command.Email, command.FirstName, command.LastName, command.ConfirmLink, h.userCounter)
	if err != nil {
		return domain.UserRegistration{}, err
	}

	if err := h.repo.Insert(ur); err != nil {
		return domain.UserRegistration{}, err
	}

	_ = h.userRegisteredHandler.Handle(UserRegisteredEvent{ur})

	return ur, nil
}
