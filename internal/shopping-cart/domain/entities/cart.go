package entities

import "github.com/google/uuid"

type CartItem struct {
    ProductID string
    Quantity  int
}

type Cart struct {
    ID     uuid.UUID
    UserID string
    Items  []CartItem
}

func NewCart(userID string) Cart {
    return Cart{
        ID:     uuid.New(),
        UserID: userID,
        Items:  []CartItem{},
    }
}

func (c *Cart) AddItem(item CartItem) {
    c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(productID string) {
    var items []CartItem
    for _, item := range c.Items {
        if item.ProductID != productID {
            items = append(items, item)
        }
    }
    c.Items = items
}
