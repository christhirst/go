-- name: CreateDemand :one
INSERT INTO
    demands
    (title, description, url, price, isDone, account_id)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;



-- name: GetDemand :one
SELECT *
FROM demands
WHERE title = $1
LIMIT 1;

-- name: ListDemands :many
SELECT *
FROM demands
ORDER BY title
LIMIT $1
OFFSET
$2;

-- name: UpdateDemand :one
UPDATE demands SET account_id = $2, description
= $3
WHERE id = $1
RETURNING *;



-- name: DeleteDemand :exec
DELETE FROM demands WHERE id = $1;

 