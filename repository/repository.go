package repository

import (
	"context"

	"github.com/golang-graphql/repository/postgres/query"
)

type Repository interface {
	CreateUser(ctx context.Context, input query.CreateUserParams) (*query.CreateUserRow, error)
	GetDetailUser(ctx context.Context, ID int32) (*query.GetUserDetailRow, error)
	GetUserByEmail(ctx context.Context, email string) (*query.GetUserByEmailRow, error)
}

type repository struct {
	qry query.Queries
}

func NewRepository(q *query.Queries) *repository {
	return &repository{qry: *q}
}
