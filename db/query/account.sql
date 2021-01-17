-- name: CreateAccount :one
INSERT INTO
    accounts (user_id, username, password)
VALUES
    ($1, $2, $3) RETURNING *;



-- name: GetAccount :one
SELECT * FROM accounts
WHERE username = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts SET password = $2
WHERE username = $1
RETURNING *;



-- name: DeleteAccount :exec
DELETE FROM accounts WHERE username = $1;