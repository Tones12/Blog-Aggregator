-- name: GetUsername :one
SELECT name FROM users WHERE id = $1;