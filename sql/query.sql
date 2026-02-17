-- name: GetCatalog :one
SELECT name, ru_name FROM catalogs WHERE id = $1 LIMIT 1;
-- name: CreateParameter :exec
INSERT INTO parameters(key, value, ru_key, ru_value) VALUES($1, $2, $3, $4);
-- name: CreateCatalog :exec
INSERT INTO catalogs(name, ru_name) VALUES($1, $2);
