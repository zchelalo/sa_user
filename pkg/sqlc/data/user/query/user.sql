-- name: CreateUser :one
INSERT INTO users (
  id,
  name,
  email,
  password,
  verified
)
VALUES ($1, $2, $3, $4, $5)
RETURNING
  id,
  name,
  email,
  verified
;

-- name: GetUser :one
SELECT
  id,
  name,
  email,
  verified
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserToAuth :one
SELECT
  id,
  name,
  email,
  password,
  verified
FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT
  id,
  name,
  email,
  verified
FROM users
WHERE verified = true
ORDER BY created_at DESC
OFFSET $1
LIMIT $2;

-- name: UpdateUser :one
UPDATE users
SET
  name = $2,
  email = $3,
  password = $4,
  verified = $5
WHERE id = $1
RETURNING
  id,
  name,
  email,
  verified
;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CountUsers :one
SELECT COUNT(*) FROM users
WHERE verified = true;