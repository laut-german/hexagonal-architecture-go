package commands

import (
	"hexagonal-architecture-go/internal/user/domain/entities"
	"hexagonal-architecture-go/internal/user/domain/ports/repositories"
)

type UpdateUserCommand struct {
	ID    string
	Name  string
	Email string
}

type UpdateUserHandler struct {
	repo repositories.UserRepository
}

func NewUpdateUserHandler(repo repositories.UserRepository) *UpdateUserHandler {
	return &UpdateUserHandler{repo: repo}
}

func (commandHandler *UpdateUserHandler) Handle(cmd UpdateUserCommand) error {
	user := entities.User{
		ID:    cmd.ID,
		Name:  cmd.Name,
		Email: cmd.Email,
	}
	return commandHandler.repo.Update(user)
}
