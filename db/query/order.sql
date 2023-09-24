-- name: CreateOrder :one
INSERT INTO orders (
  customer_id,
  service_id,
  order_status,
  order_delivered 
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE order_id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY order_id
LIMIT $1
OFFSET $2;