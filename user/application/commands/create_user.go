package commands

import (
	"hexagonal-architecture-go/user/domain/entities"
	"hexagonal-architecture-go/user/domain/repositories"
	"time"
)


type CreateUserCommand struct {
    Name  string
    Email string
}

type CreateUserHandler struct {
    repo repositories.UserRepository
}

func NewCreateUserHandler(repo repositories.UserRepository) *CreateUserHandler {
    return &CreateUserHandler{repo: repo}
}

func (commandHandler *CreateUserHandler) Handle(cmd CreateUserCommand) error {
    user := entities.User{
        Name:  cmd.Name,
        Email: cmd.Email,
        Created_At: time.Now(),
    }
    return commandHandler.repo.Save(user)
}
