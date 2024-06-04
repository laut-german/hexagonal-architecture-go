package queries

import (
	"hexagonal-architecture-go/user/domain/entities"
	"hexagonal-architecture-go/user/domain/repositories"
)


type ListUsersQuery struct {}

type ListUsersHandler struct {
    repo repositories.UserRepository
}

func NewListUsersHandler(repo repositories.UserRepository) *ListUsersHandler {
    return &ListUsersHandler{repo: repo}
}

func (queryHandler *ListUsersHandler) Handle(query ListUsersQuery) ([]entities.User, error) {
    return queryHandler.repo.List()
}
