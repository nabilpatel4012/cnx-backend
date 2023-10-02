package db

import (
	"context"
	"database/sql"
	"fmt"
)

// This Store provides all the functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function whithin a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type OrderTxParams struct {
	CustomerID  int64  `json:"customer_id"`
	ServiceID   int64  `json:"service_id"`
	OrderStatus string `json:"order_status"`
}

type OrderTxResult struct {
	Order   Order   `json:"order"`
	User    User    `json:"customer_id"`
	Service Service `json:"service_id"`
	// OrderStatus string `json:"order_status"`
}

// OrderTx performs a ordering from user to order
// It will perform full ordering process (does not imply real process) in single DB transaction
func (store *Store) OrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Update customer's order count
		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			TotalOrders: +1,
		})
		if err != nil {
			return err
		}

		// Create a new order record
		result.Order, err = q.CreateOrder(ctx, CreateOrderParams{
			CustomerID: int32(arg.CustomerID),
			ServiceID:  int32(arg.ServiceID),
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
