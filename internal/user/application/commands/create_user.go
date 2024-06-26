package commands

import (
	"hexagonal-architecture-go/internal/user/application/responses"
	"hexagonal-architecture-go/internal/user/domain/entities"
	userDomainErrors "hexagonal-architecture-go/internal/user/domain/errors"
	"hexagonal-architecture-go/internal/user/domain/ports/repositories"
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
	userExist, err := commandHandler.repo.FindByEmail(cmd.Email)
	if err != nil {
		return responses.CreateUserResponse{}, err
	}
	if userExist != nil {
		return responses.CreateUserResponse{}, &userDomainErrors.UserAlreadyExistsError{Email: cmd.Email}
	}
	user := entities.Create(cmd.Name, cmd.Email)
	savedUser, err := commandHandler.repo.Save(*user)
	if err != nil {
		return responses.CreateUserResponse{}, err
	}
	return responses.NewUserResponse(savedUser.ID, savedUser.Name, savedUser.Email), nil
}
