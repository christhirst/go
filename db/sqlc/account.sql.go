// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO
    accounts (username, password)
VALUES
    ($1, $2) RETURNING username, password
`

type CreateAccountParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Username, arg.Password)
	var i Account
	err := row.Scan(&i.Username, &i.Password)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts WHERE username = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, username)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT username, password FROM accounts
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, username string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, username)
	var i Account
	err := row.Scan(&i.Username, &i.Password)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT username, password FROM accounts
ORDER BY username
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(&i.Username, &i.Password); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts SET password = $2
WHERE username = $1
RETURNING username, password
`

type UpdateAccountParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.Username, arg.Password)
	var i Account
	err := row.Scan(&i.Username, &i.Password)
	return i, err
}
