package db

import (
	"context"
	"database/sql"
	"errors"
	"hexagonal-architecture-go/internal/shopping-cart/domain/entities"
	"hexagonal-architecture-go/internal/shopping-cart/domain/ports/repositories"

	"github.com/google/uuid"
)

type PostgresCartRepository struct {
	db *sql.DB
}

func NewPostgresCartRepository(db *sql.DB) repositories.CartRepository {
	return &PostgresCartRepository{db: db}
}

func (r *PostgresCartRepository) GetCartByUserID(ctx context.Context, userID string) (entities.Cart, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, user_id FROM carts WHERE user_id = $1", userID)

	var cart entities.Cart
	var id string

	if err := row.Scan(&id, &cart.UserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Cart{}, nil // Return an empty cart if no cart is found
		}
		return entities.Cart{}, err
	}

	cart.ID, _ = uuid.Parse(id)

	rows, err := r.db.QueryContext(ctx, "SELECT product_id, quantity FROM cart_items WHERE cart_id = $1", id)
	if err != nil {
		return entities.Cart{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entities.CartItem
		if err := rows.Scan(&item.ProductID, &item.Quantity); err != nil {
			return entities.Cart{}, err
		}
		cart.Items = append(cart.Items, item)
	}

	return cart, nil
}

func (r *PostgresCartRepository) SaveCart(ctx context.Context, cart entities.Cart) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
        INSERT INTO carts (id, user_id) VALUES ($1, $2)
        ON CONFLICT (id) DO UPDATE SET user_id = EXCLUDED.user_id
    `, cart.ID, cart.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM cart_items WHERE cart_id = $1", cart.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ($1, $2, $3)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, item := range cart.Items {
		if _, err := stmt.ExecContext(ctx, cart.ID, item.ProductID, item.Quantity); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
