package usecase

import (
	"context"
	"database/sql"

	"github.com/golang-graphql/graph/model"
	"github.com/golang-graphql/repository/postgres/query"
	"github.com/golang-graphql/utils"
)

func (uc *usecase) CreateUser(ctx context.Context, input *model.SignInParams) (*model.MainResponse, error) {

	hashPassword := utils.HashPassword(uc.saltText, input.Password)

	param := query.CreateUserParams{
		Name:     sql.NullString{String: input.Name, Valid: true},
		Email:    sql.NullString{String: input.Email, Valid: true},
		Address:  sql.NullString{String: *input.Address, Valid: true},
		Password: sql.NullString{String: hashPassword, Valid: true},
	}

	res, err := uc.repository.CreateUser(ctx, param)
	if err != nil {
		return nil, err
	}

	result := &model.MainResponse{
		IsSuccess: true,
		Message:   "Success Create Data",
		Data: &model.UserDataResponse{
			Name:    &res.Name.String,
			Email:   &res.Email.String,
			Address: &res.Address.String,
		},
		Status: 200,
	}

	return result, nil
}

func (uc *usecase) GetDetailUser(ctx context.Context, input *model.GetUserParam) (*model.MainResponseList, error) {
	return nil, nil
}
