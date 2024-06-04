package commands

import (
	"hexagonal-architecture-go/user/application/responses"
	"hexagonal-architecture-go/user/domain/entities"
	"hexagonal-architecture-go/user/domain/repositories"
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

func (commandHandler *CreateUserHandler) Handle(cmd CreateUserCommand) (responses.CreateUserResponse, error) {
	user := entities.Create(cmd.Name, cmd.Email)
	savedUser, err := commandHandler.repo.Save(*user)
	if err != nil {
		return responses.CreateUserResponse{}, err
	}
	return responses.NewUserResponse(savedUser.ID, savedUser.Name, savedUser.Email), nil
}
