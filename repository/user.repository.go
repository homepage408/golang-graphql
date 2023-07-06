package repository

import (
	"context"

	"github.com/golang-graphql/repository/postgres/query"
)

func (ar *repository) CreateUser(ctx context.Context, input query.CreateUserParams) (*query.CreateUserRow, error) {
	res, err := ar.qry.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (ar *repository) GetDetailUser(ctx context.Context, input query.GetUserDetailRow) error {
	return nil
}
