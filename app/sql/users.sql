-- name: CreateUser :execlastid
INSERT INTO users (name) VALUES (?);

-- name: GetUser :one
SELECT id, name, created_at, updated_at
FROM users
WHERE id = ? LIMIT 1;

-- name: UpdateUserName :exec
UPDATE users
SET name = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: ListUsers :many
SELECT id, name, created_at, updated_at FROM users;
