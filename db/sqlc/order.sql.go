// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: order.sql

package db

import (
	"context"
	"database/sql"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
  customer_id,
  service_id,
  order_status,
  order_delivered 
) VALUES (
  $1, $2, $3, $4
) RETURNING order_id, customer_id, service_id, order_status, order_started, order_delivered
`

type CreateOrderParams struct {
	CustomerID     int32        `json:"customer_id"`
	ServiceID      int32        `json:"service_id"`
	OrderStatus    string       `json:"order_status"`
	OrderDelivered sql.NullTime `json:"order_delivered"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.CustomerID,
		arg.ServiceID,
		arg.OrderStatus,
		arg.OrderDelivered,
	)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.CustomerID,
		&i.ServiceID,
		&i.OrderStatus,
		&i.OrderStarted,
		&i.OrderDelivered,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders WHERE order_id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, orderID int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, orderID)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT order_id, customer_id, service_id, order_status, order_started, order_delivered FROM orders
WHERE order_id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, orderID int32) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, orderID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.CustomerID,
		&i.ServiceID,
		&i.OrderStatus,
		&i.OrderStarted,
		&i.OrderDelivered,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT order_id, customer_id, service_id, order_status, order_started, order_delivered FROM orders
ORDER BY order_id
LIMIT $1
OFFSET $2
`

type ListOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.CustomerID,
			&i.ServiceID,
			&i.OrderStatus,
			&i.OrderStarted,
			&i.OrderDelivered,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :one
UPDATE orders 
SET order_status = $2
WHERE order_id = $1
RETURNING order_id, customer_id, service_id, order_status, order_started, order_delivered
`

type UpdateOrderParams struct {
	OrderID     int32  `json:"order_id"`
	OrderStatus string `json:"order_status"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrder, arg.OrderID, arg.OrderStatus)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.CustomerID,
		&i.ServiceID,
		&i.OrderStatus,
		&i.OrderStarted,
		&i.OrderDelivered,
	)
	return i, err
}
