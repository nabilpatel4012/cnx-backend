// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  phone,
  address 
) VALUES (
  $1, $2, $3, $4
) RETURNING user_id, name, email, phone, address, total_orders, created_at
`

type CreateUserParams struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Address,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.TotalOrders,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, name, email, phone, address, total_orders, created_at FROM users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.TotalOrders,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, name, email, phone, address, total_orders, created_at FROM users
ORDER BY user_id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.Address,
			&i.TotalOrders,
			&i.CreatedAt,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET email = $2,
address = $3,
phone = $4,
total_orders = $5
WHERE user_id = $1
RETURNING user_id, name, email, phone, address, total_orders, created_at
`

type UpdateUserParams struct {
	UserID      int32  `json:"user_id"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	TotalOrders int32  `json:"total_orders"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.UserID,
		arg.Email,
		arg.Address,
		arg.Phone,
		arg.TotalOrders,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.TotalOrders,
		&i.CreatedAt,
	)
	return i, err
}