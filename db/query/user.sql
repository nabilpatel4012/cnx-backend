-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  phone,
  address,
  hashed_password 
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users 
SET email = $2,
address = $3,
phone = $4,
total_orders = $5,
hashed_password = $6
WHERE user_id = $1
RETURNING *;

-- name: UpdateUserOrder :one
UPDATE users 
SET total_orders = $2
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;