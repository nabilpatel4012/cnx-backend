-- name: CreateOrder :one
INSERT INTO orders (
  customer_id,
  service_id,
  order_status,
  order_started,
  order_delivered 
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;