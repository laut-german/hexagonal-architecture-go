package responses

import "hexagonal-architecture-go/shopping-cart/domain/entities"

type CartResponse struct {
	ID     string              `json:"id"`
	UserID string              `json:"user_id"`
	Items  []CartItemResponse `json:"items"`
}

type CartItemResponse struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func NewCartResponse(cart entities.Cart) CartResponse {
	items := make([]CartItemResponse, len(cart.Items))
	for i, item := range cart.Items {
		items[i] = CartItemResponse{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}
	return CartResponse{
		ID:     cart.ID.String(),
		UserID: cart.UserID,
		Items:  items,
	}
}
