package commands

import (
	"context"
	"hexagonal-architecture-go/internal/shopping-cart/domain/entities"
	"hexagonal-architecture-go/internal/shopping-cart/domain/ports/repositories"

	"github.com/google/uuid"
)

type AddItemToCartCommand struct {
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type AddItemToCartHandler struct {
	repository repositories.CartRepository
}

func NewAddItemToCartHandler(repo repositories.CartRepository) *AddItemToCartHandler {
	return &AddItemToCartHandler{repository: repo}
}

func (commandHandler *AddItemToCartHandler) Handle(ctx context.Context, cmd AddItemToCartCommand) error {
	cart, err := commandHandler.repository.GetCartByUserID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	if cart.ID == uuid.Nil {
		cart = entities.NewCart(cmd.UserID)
	}

	cart.AddItem(entities.CartItem{
		ProductID: cmd.ProductID,
		Quantity:  cmd.Quantity,
	})

	return commandHandler.repository.SaveCart(ctx, cart)
}
