-- -- name: GetAuthor :one
-- SELECT * FROM authors WHERE id = $1 LIMIT 1;

-- DOMAINS ------------------------------------

-- name: InsertOneDomain :one
INSERT INTO domains (name) VALUES
($1) RETURNING *;

-- name: SelectOneDomainByName :one
SELECT * FROM domains WHERE name = $1 LIMIT 1;

-- NODES --------------------------------------

-- name: InsertOneNode :one
INSERT INTO nodes (domain_id) VALUES
($1)
RETURNING *;

-- name: InsertManyNodes :exec
INSERT INTO nodes (domain_id) VALUES (
    unnest(@a_array::text[])
) RETURNING *;

-- name: SelectManyNodesWithPagination :many
SELECT * FROM nodes
WHERE id > $1 
ORDER BY id ASC
LIMIT $2;