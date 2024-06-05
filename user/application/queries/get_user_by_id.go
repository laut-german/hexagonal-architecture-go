package queries

import (
	"hexagonal-architecture-go/user/domain/entities"
	"hexagonal-architecture-go/user/domain/ports/repositories"
)

type GetUserByIDQuery struct {
	ID string
}

type GetUserByIDHandler struct {
	repo repositories.UserRepository
}

func NewGetUserByIDHandler(repo repositories.UserRepository) *GetUserByIDHandler {
	return &GetUserByIDHandler{repo: repo}
}

func (queryHandler *GetUserByIDHandler) Handle(query GetUserByIDQuery) (*entities.User, error) {
	return queryHandler.repo.FindByID(query.ID)
}
