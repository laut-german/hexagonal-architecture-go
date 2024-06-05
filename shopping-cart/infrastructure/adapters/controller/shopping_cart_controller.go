package controller

import (
    "encoding/json"
    "net/http"
    "hexagonal-architecture-go/shopping-cart/application/commands"
    "hexagonal-architecture-go/shopping-cart/application/queries"
    "hexagonal-architecture-go/shopping-cart/application/responses"
    "hexagonal-architecture-go/shopping-cart/domain/ports/queue"
    "hexagonal-architecture-go/shopping-cart/domain/ports/repositories"
)

type ShoppingCartController struct {
    addItemHandler         *commands.AddItemToCartHandler
    removeItemHandler      *commands.RemoveItemFromCartHandler
    getCartByUserIDHandler *queries.GetCartByUserIDHandler
    listCartItemsHandler   *queries.ListCartItemsHandler
    queueProducer          queue.QueueProducer
}

func NewShoppingCartController(cartRepo repositories.CartRepository, queueProducer queue.QueueProducer) *ShoppingCartController {
    return &ShoppingCartController{
        addItemHandler:         commands.NewAddItemToCartHandler(cartRepo),
        removeItemHandler:      commands.NewRemoveItemFromCartHandler(cartRepo),
        getCartByUserIDHandler: queries.NewGetCartByUserIDHandler(cartRepo),
        listCartItemsHandler:   queries.NewListCartItemsHandler(cartRepo),
        queueProducer:          queueProducer,
    }
}

func (c *ShoppingCartController) AddItemToCart(w http.ResponseWriter, r *http.Request) {
    var cmd commands.AddItemToCartCommand
    if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    message, err := json.Marshal(cmd)
    if err != nil {
        http.Error(w, "Failed to marshal message", http.StatusInternalServerError)
        return
    }

    if err := c.queueProducer.Publish(r.Context(), message); err != nil {
        http.Error(w, "Failed to publish message", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (c *ShoppingCartController) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
    var cmd commands.RemoveItemFromCartCommand
    if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    message, err := json.Marshal(cmd)
    if err != nil {
        http.Error(w, "Failed to marshal message", http.StatusInternalServerError)
        return
    }

    if err := c.queueProducer.Publish(r.Context(), message); err != nil {
        http.Error(w, "Failed to publish message", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (c *ShoppingCartController) GetCartByUserID(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user_id")
    if userID == "" {
        http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
        return
    }

    cart, err := c.getCartByUserIDHandler.Handle(r.Context(), queries.GetCartByUserIDQuery{UserID: userID})
    if err != nil {
        http.Error(w, "Failed to get cart", http.StatusInternalServerError)
        return
    }

    response := responses.NewCartResponse(cart)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (c *ShoppingCartController) ListCartItems(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user_id")
    if userID == "" {
        http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
        return
    }

    items, err := c.listCartItemsHandler.Handle(r.Context(), queries.ListCartItemsQuery{UserID: userID})
    if err != nil {
        http.Error(w, "Failed to list cart items", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}
