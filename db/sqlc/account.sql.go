// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts
(user_id, category_id, title, type, description, value, date)
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_id, category_id, title, type, description, value, date, created_at
`

type CreateAccountParams struct {
	UserID      int32     `json:"user_id"`
	CategoryID  int32     `json:"category_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Value       int32     `json:"value"`
	Date        time.Time `json:"date"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.UserID,
		arg.CategoryID,
		arg.Title,
		arg.Type,
		arg.Description,
		arg.Value,
		arg.Date,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
 WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccounts = `-- name: GetAccounts :many
SELECT a.id,
       a.user_id,
       a.category_id,
       a.title,
       a.type,
       a.description,
       a.value,
       a.date,
       a.created_at,
       c.title AS category_title
  FROM accounts a
  LEFT JOIN categories c ON c.id = a.category_id
 WHERE a.user_id = $1
   AND a.type = $2
   AND a.category_id = $3
   AND a.title like $4
   AND a.description like $5
   AND a.date = $6
`

type GetAccountsParams struct {
	UserID      int32     `json:"user_id"`
	Type        string    `json:"type"`
	CategoryID  int32     `json:"category_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type GetAccountsRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	CategoryID    int32          `json:"category_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccounts,
		arg.UserID,
		arg.Type,
		arg.CategoryID,
		arg.Title,
		arg.Description,
		arg.Date,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsRow{}
	for rows.Next() {
		var i GetAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
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

const getAccountsById = `-- name: GetAccountsById :one
SELECT id, user_id, category_id, title, type, description, value, date, created_at
FROM accounts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAccountsById(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountsById, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountsByUserIdAndType = `-- name: GetAccountsByUserIdAndType :many
SELECT a.id,
       a.user_id,
       a.category_id,
       a.title,
       a.type,
       a.description,
       a.value,
       a.date,
       a.created_at,
       c.title AS category_title
 FROM accounts a
 LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = $1
  AND a.type = $2
`

type GetAccountsByUserIdAndTypeParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

type GetAccountsByUserIdAndTypeRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	CategoryID    int32          `json:"category_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccountsByUserIdAndType(ctx context.Context, arg GetAccountsByUserIdAndTypeParams) ([]GetAccountsByUserIdAndTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndType, arg.UserID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndTypeRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndTypeRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
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

const getAccountsByUserIdAndTypeAndCategoryId = `-- name: GetAccountsByUserIdAndTypeAndCategoryId :many
SELECT a.id,
       a.user_id,
       a.category_id,
       a.title,
       a.type,
       a.description,
       a.value,
       a.date,
       a.created_at,
       c.title AS category_title
 FROM accounts a
 LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = $1
  AND a.type = $2
  AND a.category_id = $3
`

type GetAccountsByUserIdAndTypeAndCategoryIdParams struct {
	UserID     int32  `json:"user_id"`
	Type       string `json:"type"`
	CategoryID int32  `json:"category_id"`
}

type GetAccountsByUserIdAndTypeAndCategoryIdRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	CategoryID    int32          `json:"category_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccountsByUserIdAndTypeAndCategoryId(ctx context.Context, arg GetAccountsByUserIdAndTypeAndCategoryIdParams) ([]GetAccountsByUserIdAndTypeAndCategoryIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndTypeAndCategoryId, arg.UserID, arg.Type, arg.CategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndTypeAndCategoryIdRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndTypeAndCategoryIdRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
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

const getAccountsByUserIdAndTypeAndCategoryIdAndTitle = `-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTitle :many
SELECT a.id,
       a.user_id,
       a.category_id,
       a.title,
       a.type,
       a.description,
       a.value,
       a.date,
       a.created_at,
       c.title AS category_title
 FROM accounts a
 LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = $1
  AND a.type = $2
  AND a.category_id = $3
  AND a.title like $4
`

type GetAccountsByUserIdAndTypeAndCategoryIdAndTitleParams struct {
	UserID     int32  `json:"user_id"`
	Type       string `json:"type"`
	CategoryID int32  `json:"category_id"`
	Title      string `json:"title"`
}

type GetAccountsByUserIdAndTypeAndCategoryIdAndTitleRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	CategoryID    int32          `json:"category_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccountsByUserIdAndTypeAndCategoryIdAndTitle(ctx context.Context, arg GetAccountsByUserIdAndTypeAndCategoryIdAndTitleParams) ([]GetAccountsByUserIdAndTypeAndCategoryIdAndTitleRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndTypeAndCategoryIdAndTitle,
		arg.UserID,
		arg.Type,
		arg.CategoryID,
		arg.Title,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndTypeAndCategoryIdAndTitleRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndTypeAndCategoryIdAndTitleRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
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

const getAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription = `-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription :many
SELECT a.id,
       a.user_id,
       a.category_id,
       a.title,
       a.type,
       a.description,
       a.value,
       a.date,
       a.created_at,
       c.title AS category_title
 FROM accounts a
 LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = $1
  AND a.type = $2
  AND a.category_id = $3
  AND a.title like $4
  AND a.description like $5
`

type GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescriptionParams struct {
	UserID      int32  `json:"user_id"`
	Type        string `json:"type"`
	CategoryID  int32  `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescriptionRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	CategoryID    int32          `json:"category_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription(ctx context.Context, arg GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescriptionParams) ([]GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescriptionRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription,
		arg.UserID,
		arg.Type,
		arg.CategoryID,
		arg.Title,
		arg.Description,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescriptionRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescriptionRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
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

const getAccountsByUserIdAndTypeAndDate = `-- name: GetAccountsByUserIdAndTypeAndDate :many
SELECT a.id,
       a.user_id,
       a.category_id,
       a.title,
       a.type,
       a.description,
       a.value,
       a.date,
       a.created_at,
       c.title AS category_title
 FROM accounts a
 LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = $1
  AND a.type = $2
  AND a.date = $3
`

type GetAccountsByUserIdAndTypeAndDateParams struct {
	UserID int32     `json:"user_id"`
	Type   string    `json:"type"`
	Date   time.Time `json:"date"`
}

type GetAccountsByUserIdAndTypeAndDateRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	CategoryID    int32          `json:"category_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccountsByUserIdAndTypeAndDate(ctx context.Context, arg GetAccountsByUserIdAndTypeAndDateParams) ([]GetAccountsByUserIdAndTypeAndDateRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndTypeAndDate, arg.UserID, arg.Type, arg.Date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndTypeAndDateRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndTypeAndDateRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
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

const getAccountsByUserIdAndTypeAndDescription = `-- name: GetAccountsByUserIdAndTypeAndDescription :many
SELECT a.id,
       a.user_id,
       a.category_id,
       a.title,
       a.type,
       a.description,
       a.value,
       a.date,
       a.created_at,
       c.title AS category_title
 FROM accounts a
 LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = $1
  AND a.type = $2
  AND a.description like $3
`

type GetAccountsByUserIdAndTypeAndDescriptionParams struct {
	UserID      int32  `json:"user_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type GetAccountsByUserIdAndTypeAndDescriptionRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	CategoryID    int32          `json:"category_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccountsByUserIdAndTypeAndDescription(ctx context.Context, arg GetAccountsByUserIdAndTypeAndDescriptionParams) ([]GetAccountsByUserIdAndTypeAndDescriptionRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndTypeAndDescription, arg.UserID, arg.Type, arg.Description)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndTypeAndDescriptionRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndTypeAndDescriptionRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
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

const getAccountsByUserIdAndTypeAndTitle = `-- name: GetAccountsByUserIdAndTypeAndTitle :many
SELECT a.id,
       a.user_id,
       a.category_id,
       a.title,
       a.type,
       a.description,
       a.value,
       a.date,
       a.created_at,
       c.title AS category_title
 FROM accounts a
 LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = $1
  AND a.type = $2
  AND a.title like $3
`

type GetAccountsByUserIdAndTypeAndTitleParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
	Title  string `json:"title"`
}

type GetAccountsByUserIdAndTypeAndTitleRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	CategoryID    int32          `json:"category_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccountsByUserIdAndTypeAndTitle(ctx context.Context, arg GetAccountsByUserIdAndTypeAndTitleParams) ([]GetAccountsByUserIdAndTypeAndTitleRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndTypeAndTitle, arg.UserID, arg.Type, arg.Title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndTypeAndTitleRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndTypeAndTitleRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
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

const getAccountsGraph = `-- name: GetAccountsGraph :one
SELECT COUNT(*)
  FROM accounts
 WHERE user_id = $1
   AND type = $2
`

type GetAccountsGraphParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetAccountsGraph(ctx context.Context, arg GetAccountsGraphParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountsGraph, arg.UserID, arg.Type)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAccountsReports = `-- name: GetAccountsReports :one
SELECT SUM(value) AS sum_value
  FROM accounts
 WHERE user_id = $1
   AND type = $2
`

type GetAccountsReportsParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetAccountsReports(ctx context.Context, arg GetAccountsReportsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountsReports, arg.UserID, arg.Type)
	var sum_value int64
	err := row.Scan(&sum_value)
	return sum_value, err
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
   SET title = $2,
       description = $3,
       value = $4
 WHERE id = $1 RETURNING id, user_id, category_id, title, type, description, value, date, created_at
`

type UpdateAccountParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Value,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}
