package repository

import (
	"context"

	"github.com/golang-graphql/repository/postgres/query"
)

type Repository interface {
	CreateUser(ctx context.Context, input query.CreateUserParams) error
	GetDetailUser(ctx context.Context, input query.GetUserDetailRow) error
}

type repository struct {
}

func NewRepository(ctx context.Context)
