-- name: CreateService :one
INSERT INTO services (
  service_name,
  service_price 
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetService :one
SELECT * FROM services
WHERE service_id = $1 LIMIT 1;

-- name: ListLimitedServices :many
SELECT * FROM services
ORDER BY service_id
LIMIT $1
OFFSET $2;

-- name: ListAllServices :many
SELECT * FROM services
ORDER BY service_id;

-- name: UpdateService :one
UPDATE services 
SET service_name = $2,
service_price = $3
WHERE service_id = $1
RETURNING *;