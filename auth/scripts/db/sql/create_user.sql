-- name: InsertUser :one
INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING uid;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1;