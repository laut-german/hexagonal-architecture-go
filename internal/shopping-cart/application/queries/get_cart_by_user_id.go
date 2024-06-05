package queries

import (
	"context"
	"hexagonal-architecture-go/internal/shopping-cart/domain/entities"
	"hexagonal-architecture-go/internal/shopping-cart/domain/ports/repositories"
)

type GetCartByUserIDQuery struct {
	UserID string `json:"user_id"`
}

type GetCartByUserIDHandler struct {
	repository repositories.CartRepository
}

func NewGetCartByUserIDHandler(repo repositories.CartRepository) *GetCartByUserIDHandler {
	return &GetCartByUserIDHandler{repository: repo}
}

func (queryHandler *GetCartByUserIDHandler) Handle(ctx context.Context, query GetCartByUserIDQuery) (entities.Cart, error) {
	return queryHandler.repository.GetCartByUserID(ctx, query.UserID)
}
