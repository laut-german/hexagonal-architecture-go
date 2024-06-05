package queries

import (
	"context"
	"hexagonal-architecture-go/internal/shopping-cart/domain/entities"
	"hexagonal-architecture-go/internal/shopping-cart/domain/ports/repositories"
)

type ListCartItemsQuery struct {
	UserID string `json:"user_id"`
}

type ListCartItemsHandler struct {
	repository repositories.CartRepository
}

func NewListCartItemsHandler(repo repositories.CartRepository) *ListCartItemsHandler {
	return &ListCartItemsHandler{repository: repo}
}

func (queryHandler *ListCartItemsHandler) Handle(ctx context.Context, query ListCartItemsQuery) ([]entities.CartItem, error) {
	cart, err := queryHandler.repository.GetCartByUserID(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	return cart.Items, nil
}
