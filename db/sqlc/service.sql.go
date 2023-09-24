// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: service.sql

package db

import (
	"context"
)

const createService = `-- name: CreateService :one
INSERT INTO services (
  service_name,
  service_price 
) VALUES (
  $1, $2
) RETURNING service_id, service_name, service_price
`

type CreateServiceParams struct {
	ServiceName  string `json:"service_name"`
	ServicePrice int64  `json:"service_price"`
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) (Service, error) {
	row := q.db.QueryRowContext(ctx, createService, arg.ServiceName, arg.ServicePrice)
	var i Service
	err := row.Scan(&i.ServiceID, &i.ServiceName, &i.ServicePrice)
	return i, err
}

const deleteService = `-- name: DeleteService :exec
DELETE FROM services WHERE service_id = $1
`

func (q *Queries) DeleteService(ctx context.Context, serviceID int32) error {
	_, err := q.db.ExecContext(ctx, deleteService, serviceID)
	return err
}

const getService = `-- name: GetService :one
SELECT service_id, service_name, service_price FROM services
WHERE service_id = $1 LIMIT 1
`

func (q *Queries) GetService(ctx context.Context, serviceID int32) (Service, error) {
	row := q.db.QueryRowContext(ctx, getService, serviceID)
	var i Service
	err := row.Scan(&i.ServiceID, &i.ServiceName, &i.ServicePrice)
	return i, err
}

const listAllServices = `-- name: ListAllServices :many
SELECT service_id, service_name, service_price FROM services
ORDER BY service_id
`

func (q *Queries) ListAllServices(ctx context.Context) ([]Service, error) {
	rows, err := q.db.QueryContext(ctx, listAllServices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Service
	for rows.Next() {
		var i Service
		if err := rows.Scan(&i.ServiceID, &i.ServiceName, &i.ServicePrice); err != nil {
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

const listLimitedServices = `-- name: ListLimitedServices :many
SELECT service_id, service_name, service_price FROM services
ORDER BY service_id
LIMIT $1
OFFSET $2
`

type ListLimitedServicesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListLimitedServices(ctx context.Context, arg ListLimitedServicesParams) ([]Service, error) {
	rows, err := q.db.QueryContext(ctx, listLimitedServices, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Service
	for rows.Next() {
		var i Service
		if err := rows.Scan(&i.ServiceID, &i.ServiceName, &i.ServicePrice); err != nil {
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

const updateService = `-- name: UpdateService :one
UPDATE services 
SET service_name = $2,
service_price = $3
WHERE service_id = $1
RETURNING service_id, service_name, service_price
`

type UpdateServiceParams struct {
	ServiceID    int32  `json:"service_id"`
	ServiceName  string `json:"service_name"`
	ServicePrice int64  `json:"service_price"`
}

func (q *Queries) UpdateService(ctx context.Context, arg UpdateServiceParams) (Service, error) {
	row := q.db.QueryRowContext(ctx, updateService, arg.ServiceID, arg.ServiceName, arg.ServicePrice)
	var i Service
	err := row.Scan(&i.ServiceID, &i.ServiceName, &i.ServicePrice)
	return i, err
}
