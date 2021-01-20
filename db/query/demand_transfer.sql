-- name: CreateDemand_transfer :one
INSERT INTO demand_transfer (
  from_account_id,
  to_account_id,
  demand_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetDemand_transfer :one
SELECT * FROM demand_transfer
WHERE id = $1 LIMIT 1;

-- name: ListDemand_transfer :many
SELECT * FROM demand_transfer
WHERE 
    from_account_id = $1 OR
    to_account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: DeleteDemand_transfer :exec
DELETE FROM demand_transfer WHERE id = $1;