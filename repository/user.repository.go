package repository

import (
	"context"
	"database/sql"

	"github.com/golang-graphql/repository/postgres/query"
)

func (ar *repository) CreateUser(ctx context.Context, input query.CreateUserParams) (*query.CreateUserRow, error) {
	res, err := ar.qry.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (ar *repository) GetDetailUser(ctx context.Context, ID int32) (*query.GetUserDetailRow, error) {
	res, err := ar.qry.GetUserDetail(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (ar *repository) GetUserByEmail(ctx context.Context, email string) (*query.GetUserByEmailRow, error) {
	res, err := ar.qry.GetUserByEmail(ctx, sql.NullString{String: email, Valid: true})
	if err != nil {
		return nil, err
	}

	return &res, nil
}
