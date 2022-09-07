// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int32) error
	DeleteCategory(ctx context.Context, id int32) error
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error)
	GetAccountsById(ctx context.Context, id int32) (Account, error)
	GetAccountsGraph(ctx context.Context, arg GetAccountsGraphParams) (int64, error)
	GetAccountsReports(ctx context.Context, arg GetAccountsReportsParams) (int64, error)
	GetCategoriesById(ctx context.Context, id int32) (Category, error)
	GetCategoriesByUserId(ctx context.Context, arg GetCategoriesByUserIdParams) ([]Category, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
}

var _ Querier = (*Queries)(nil)
