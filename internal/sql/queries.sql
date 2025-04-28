-- name: Create :exec
INSERT INTO schools (id, name, register, unit, address, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: FindByRegister :one
SELECT * FROM schools WHERE register = $1 LIMIT 1;

-- name: FindByID :one
SELECT * FROM schools WHERE id = $1 LIMIT 1;

-- name: Update :exec
UPDATE schools SET name = $2, register = $3, unit = $4, address = $5, updated_at = $6 WHERE id = $1;

-- name: Delete :exec
DELETE FROM schools WHERE id = $1;
