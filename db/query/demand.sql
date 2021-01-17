-- name: CreateDemand :one
INSERT INTO
    demands
    (id, title, description, url, price, isDone, owner)
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
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
UPDATE demands SET owner = $2, description
= $3
WHERE title = $1
RETURNING *;



-- name: DeleteDemand :exec
DELETE FROM demands WHERE id = $1;

 