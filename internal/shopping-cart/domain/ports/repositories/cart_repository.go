package repositories

import (
	"context"
	"hexagonal-architecture-go/internal/shopping-cart/domain/entities"
)

type CartRepository interface {
	GetCartByUserID(ctx context.Context, userID string) (entities.Cart, error)
	SaveCart(ctx context.Context, cart entities.Cart) error
}
