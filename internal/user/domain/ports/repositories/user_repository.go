package repositories

import (
	"hexagonal-architecture-go/internal/user/domain/entities"
)

type UserRepository interface {
	Save(user entities.User) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	List() ([]entities.User, error)
	Update(user entities.User) error
}
