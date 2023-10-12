-- name: CreateService :one
INSERT INTO services (
  service_name,
  service_price,
  service_image 
) VALUES (
  $1, $2, $3
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
service_price = $3,
service_image = $4
WHERE service_id = $1
RETURNING *;

-- name: DeleteService :exec
DELETE FROM services WHERE service_id = $1;