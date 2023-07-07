package usecase

import (
	"context"
	"database/sql"
	"strings"

	"github.com/golang-graphql/graph/model"
	"github.com/golang-graphql/repository/postgres/query"
	"github.com/golang-graphql/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (uc *usecase) CreateUser(ctx context.Context, input *model.SignInParams) (*model.MainResponse, error) {

	toLowerEmail := strings.ToLower(input.Email)

	// Check if the user already exists or not
	_, err := uc.repository.GetUserByEmail(ctx, toLowerEmail)
	if err == nil {
		return nil, gqlerror.Errorf("email is already exists")
	}

	hashPassword := utils.HashPassword(uc.saltText, input.Password)
	param := query.CreateUserParams{
		Name:     sql.NullString{String: input.Name, Valid: true},
		Email:    sql.NullString{String: toLowerEmail, Valid: true},
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

func (uc *usecase) GetDetailUser(ctx context.Context, input int32) (*model.MainResponse, error) {
	// var name = "teguh"
	res, err := uc.repository.GetDetailUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.MainResponse{
		IsSuccess: true,
		Message:   "Success",
		Data: &model.UserDataResponse{
			// ID:      res.ID,
			Name:    &res.Name.String,
			Email:   &res.Email.String,
			Address: &res.Address.String,
		},
	}, err
}
