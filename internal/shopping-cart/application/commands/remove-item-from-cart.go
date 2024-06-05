package commands

import (
	"context"
	"hexagonal-architecture-go/internal/shopping-cart/domain/ports/repositories"

	"github.com/google/uuid"
)

type RemoveItemFromCartCommand struct {
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
}

type RemoveItemFromCartHandler struct {
	repository repositories.CartRepository
}

func NewRemoveItemFromCartHandler(repo repositories.CartRepository) *RemoveItemFromCartHandler {
	return &RemoveItemFromCartHandler{repository: repo}
}

func (commandHandler *RemoveItemFromCartHandler) Handle(ctx context.Context, cmd RemoveItemFromCartCommand) error {
	cart, err := commandHandler.repository.GetCartByUserID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	if cart.ID == uuid.Nil {
		return nil // No need to remove from a non-existing cart
	}

	cart.RemoveItem(cmd.ProductID)

	return commandHandler.repository.SaveCart(ctx, cart)
}
